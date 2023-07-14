package performance_test

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/advanced_http"
	"go.uber.org/zap"

	"github.com/solo-io/licensing/pkg/model"
	license2 "github.com/solo-io/solo-projects/pkg/license"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/setup"

	"github.com/solo-io/gloo/test/ginkgo/decorators"

	"github.com/solo-io/go-utils/contextutils"

	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/gloo/test/ginkgo/labels"

	"github.com/onsi/gomega/types"
	"github.com/solo-io/gloo/test/gomega/matchers"

	"github.com/golang/mock/gomock"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	mock_consul "github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul/mocks"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	gloohelpers "github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	. "github.com/solo-io/gloo/projects/gloo/pkg/translator"

	"time"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

// Tests are run as part of the "Nightly" action in a GHA using the default Linux runner
// More info on that machine can be found here: https://docs.github.com/en/actions/using-github-hosted-runners/about-github-hosted-runners#supported-runners-and-hardware-resources
// When developing new tests, users should manually run that action in order to test performance under the same parameters
// Results can then be found in the logs for that instance of the action
var _ = Describe("Translation - Benchmarking Tests", decorators.Performance, Label(labels.Performance), func() {
	var (
		ctrl       *gomock.Controller
		settings   *v1.Settings
		translator Translator
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		settings = &v1.Settings{}
		memoryClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		opts := bootstrap.Opts{
			Settings:  settings,
			Secrets:   memoryClientFactory,
			Upstreams: memoryClientFactory,
			Consul: bootstrap.Consul{
				ConsulWatcher: mock_consul.NewMockConsulWatcher(ctrl), // just needed to activate the consul plugin
			},
		}

		apiEmitterChan := make(chan struct{})

		// For testing purposes, load the LicensedFeatureProvider with a valid license
		licensedFeatureProvider := license2.NewLicensedFeatureProvider()
		licensedFeatureProvider.SetValidatedLicense(&license2.ValidatedLicense{
			License: &model.License{
				IssuedAt:      time.Now(),
				ExpiresAt:     time.Now(),
				RandomPayload: "",
				LicenseType:   model.LicenseType_Enterprise,
				Product:       model.Product_Gloo,
				AddOns:        nil,
			},
			Warn: nil,
			Err:  nil,
		})

		pluginRegistryFactory := setup.GetPluginRegistryFactory(opts, apiEmitterChan, licensedFeatureProvider)

		translator = NewTranslatorWithHasher(glooutils.NewSslConfigTranslator(), settings, pluginRegistryFactory(context.Background()), EnvoyCacheResourcesListToFnvHash)
	})

	// The Benchmark table takes entries consisting of an ApiSnapshot, benchmarkConfig, and labels
	// We measure the duration of the translation of the snapshot, benchmarking according to the benchmarkConfig
	// Labels are used to add context to the entry description
	DescribeTable("Benchmark table",
		func(snapBuilder *gloohelpers.ScaledSnapshotBuilder, config *gloohelpers.BenchmarkConfig, labels ...string) {
			var (
				apiSnap *v1snap.ApiSnapshot
				proxy   *v1.Proxy

				snap   cache.Snapshot
				errs   reporter.ResourceReports
				report *validation.ProxyReport

				tooFastWarningCount int
			)

			originalLogLevel := contextutils.GetLogLevel()
			contextutils.SetLogLevel(zap.ErrorLevel)
			defer contextutils.SetLogLevel(originalLogLevel)

			apiSnap = snapBuilder.Build()

			params := plugins.Params{
				Ctx:      context.Background(),
				Snapshot: apiSnap,
			}

			Expect(apiSnap.Proxies).NotTo(BeEmpty())
			proxy = apiSnap.Proxies[0]

			desc := gloohelpers.GenerateBenchmarkDesc(snapBuilder, config, labels...)

			experiment := gmeasure.NewExperiment(fmt.Sprintf("Experiment - %s", desc))

			AddReportEntry(experiment.Name, experiment)

			experiment.Sample(func(idx int) {

				// Time translation
				res, ignore, err := gloohelpers.MeasureIgnore0ns(func() {
					snap, errs, report = translator.Translate(params, proxy)
				})
				Expect(err).NotTo(HaveOccurred())

				if idx == 0 {
					// Assert expected results on the first sample
					Expect(errs.Validate()).NotTo(HaveOccurred())
					Expect(snap).NotTo(BeNil())
					Expect(report).To(Equal(validationutils.MakeReport(proxy)))
				}

				if ignore {
					tooFastWarningCount++
					return
				}

				// Benchmark total time spent
				// If desired, a field can be added to benchmarkConfig to allow benchmarking according to User and/or
				// System time
				experiment.RecordDuration(desc, res.Total)
			}, gmeasure.SamplingConfig{N: config.Iterations, Duration: config.MaxDur})

			if tooFastWarningCount > 0 {
				logger := contextutils.LoggerFrom(params.Ctx)
				logger.Warnf("entry %s registered %d 0ns measurements; consider increasing the scale of the proxy being tested for more accurate results", desc, tooFastWarningCount)
			}

			durations := experiment.Get(desc).Durations

			Expect(durations).Should(And(config.GetMatchers()...))
		},
		gloohelpers.GenerateBenchmarkDesc, // generate descriptions for table entries with nil descriptions
		Entry("basic", gloohelpers.NewInjectedSnapshotBuilder(basicSnap), basicConfig),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).WithEndpointCount(1), basicConfig, "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(1000).WithEndpointCount(1), oneKUpstreamsConfig, "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(1).WithEndpointCount(10), basicConfig, "endpoint scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(1).WithEndpointCount(1000), basicConfig, "endpoint scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).WithEndpointCount(10), basicConfig, "endpoint scale", "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).WithEndpointCount(1).
			WithUpstreamBuilder(singleHealthCheckUsBuilder), basicConfig, "single health check", "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(1000).WithEndpointCount(1).
			WithUpstreamBuilder(singleHealthCheckUsBuilder), oneKUpstreamsConfig, "single health check", "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).WithEndpointCount(1).
			WithUpstreamBuilder(multiHealthCheckUsBuilder), basicConfig, "multi health check", "upstream scale"),
		Entry(nil, gloohelpers.NewScaledSnapshotBuilder().WithUpstreamCount(1000).WithEndpointCount(1).
			WithUpstreamBuilder(multiHealthCheckUsBuilder), oneKUpstreamsConfig, "multi health check", "upstream scale"),
	)
})

// Test assets: Add blocks for logical groupings of tests, including:
// - in-line snapshot definitions for tests that require granularly-configured/heterogeneous resources (ie testing a particular field or feature)
// - benchmarkConfigs for particular groups of use cases depending on processing time requirements/expectations

/* Basic Tests */
var basicSnap = &v1snap.ApiSnapshot{
	Proxies: []*v1.Proxy{
		{
			Listeners: []*v1.Listener{
				gloohelpers.HttpListener(1),
			},
		},
	},
	Endpoints: []*v1.Endpoint{gloohelpers.Endpoint(0)},
	Upstreams: []*v1.Upstream{gloohelpers.Upstream(0)},
}

var basicConfig = &gloohelpers.BenchmarkConfig{
	Iterations: 1000,
	MaxDur:     10 * time.Second,
	GhaMatchers: []types.GomegaMatcher{
		matchers.HaveMedianLessThan(50 * time.Millisecond),
		matchers.HavePercentileLessThan(90, 100*time.Millisecond),
	},
}

/* 1k Upstreams Scale Test */
var oneKUpstreamsConfig = &gloohelpers.BenchmarkConfig{
	Iterations: 100,
	MaxDur:     30 * time.Second,
	GhaMatchers: []types.GomegaMatcher{
		matchers.HaveMedianLessThan(time.Second),
		matchers.HavePercentileLessThan(90, 2*time.Second),
	},
}

/* Advanced HTTP Plugin (Healthcheck) Test */
var (
	healthCheck = &core.HealthCheck{
		HealthyThreshold:   &wrappers.UInt32Value{Value: 3},
		UnhealthyThreshold: &wrappers.UInt32Value{Value: 3},
		HealthChecker: &core.HealthCheck_HttpHealthCheck_{
			HttpHealthCheck: &core.HealthCheck_HttpHealthCheck{
				Path: "/foo",
				ResponseAssertions: &advanced_http.ResponseAssertions{
					ResponseMatchers: []*advanced_http.ResponseMatcher{
						{
							ResponseMatch: &advanced_http.ResponseMatch{
								JsonKey: &advanced_http.JsonKey{
									Path: []*advanced_http.JsonKey_PathSegment{
										{
											Segment: &advanced_http.JsonKey_PathSegment_Key{
												Key: "key1",
											},
										},
										{
											Segment: &advanced_http.JsonKey_PathSegment_Key{
												Key: "key2",
											},
										},
									},
								},
								Source: &advanced_http.ResponseMatch_Header{
									Header: "header1",
								},
								Regex: "pattern",
							},
						},
					},
				},
			},
		},
	}

	singleHealthCheckUsBuilder = gloohelpers.NewUpstreamBuilder().WithHealthChecks([]*core.HealthCheck{healthCheck})
	multiHealthCheckUsBuilder  = gloohelpers.NewUpstreamBuilder().WithHealthChecks([]*core.HealthCheck{healthCheck, healthCheck, healthCheck, healthCheck, healthCheck})
)
