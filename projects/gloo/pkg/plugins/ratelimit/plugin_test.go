package ratelimit_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/ratelimit"

	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyratelimit "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/rate_limit/v2"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	rlconfig "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
	"github.com/gogo/protobuf/types"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	ratelimitpb "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/plugins/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/util"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var _ = Describe("Plugin", func() {
	var (
		rlSettings *ratelimitpb.Settings
		initParams plugins.InitParams
		params     plugins.Params
		rlPlugin   *Plugin
		extensions *gloov1.Extensions
		ref        core.ResourceRef
	)

	BeforeEach(func() {
		rlPlugin = NewPlugin()
		ref = core.ResourceRef{
			Name:      "test",
			Namespace: "test",
		}

		rlSettings = &ratelimitpb.Settings{
			RatelimitServerRef: &ref,
		}
		initParams = plugins.InitParams{}
	})

	JustBeforeEach(func() {
		settingsStruct, err := util.MessageToStruct(rlSettings)
		Expect(err).NotTo(HaveOccurred())

		glooExtensions := map[string]*types.Struct{
			ExtensionName: settingsStruct,
		}
		extensions = &gloov1.Extensions{
			Configs: glooExtensions,
		}
		initParams.ExtensionsSettings = extensions
		rlPlugin.Init(initParams)
	})

	It("should fave fail mode deny off by default", func() {

		filters, err := rlPlugin.HttpFilters(params, nil)
		Expect(err).NotTo(HaveOccurred())

		Expect(filters).To(HaveLen(2))
		for _, f := range filters {
			cfg := getConfig(f.HttpFilter)
			Expect(cfg.FailureModeDeny).To(BeFalse())
		}

		hundredms := time.Millisecond * 100
		expectedConfig := &envoyratelimit.RateLimit{
			Domain:          "ingress",
			FailureModeDeny: false,
			Stage:           0,
			Timeout:         &hundredms,
			RequestType:     "both",
			RateLimitService: &rlconfig.RateLimitServiceConfig{
				GrpcService: &envoycore.GrpcService{TargetSpecifier: &envoycore.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &envoycore.GrpcService_EnvoyGrpc{
						ClusterName: translator.UpstreamToClusterName(ref),
					},
				}},
			},
		}

		cfg1 := getConfig(filters[1].HttpFilter)
		Expect(cfg1).To(BeEquivalentTo(expectedConfig))

		expectedConfig.Domain = "custom"
		expectedConfig.Stage = 1

		cfg2 := getConfig(filters[0].HttpFilter)
		Expect(cfg2).To(BeEquivalentTo(expectedConfig))

	})

	It("default timeout is 100ms", func() {
		filters, err := rlPlugin.HttpFilters(params, nil)
		Expect(err).NotTo(HaveOccurred())
		timeout := time.Millisecond * 100
		Expect(filters).To(HaveLen(2))
		for _, f := range filters {
			cfg := getConfig(f.HttpFilter)
			Expect(*cfg.Timeout).To(Equal(timeout))
		}
	})

	Context("fail mode deny", func() {

		BeforeEach(func() {
			rlSettings.DenyOnFail = true
		})

		It("should fave fail mode deny on", func() {
			filters, err := rlPlugin.HttpFilters(params, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(filters).To(HaveLen(2))
			for _, f := range filters {
				cfg := getConfig(f.HttpFilter)
				Expect(cfg.FailureModeDeny).To(BeTrue())
			}
		})
	})

	Context("timeout", func() {

		BeforeEach(func() {
			s := time.Second
			rlSettings.RequestTimeout = &s
		})

		It("should custom timeout set", func() {
			filters, err := rlPlugin.HttpFilters(params, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(filters).To(HaveLen(2))
			for _, f := range filters {
				cfg := getConfig(f.HttpFilter)
				Expect(*cfg.Timeout).To(Equal(time.Second))
			}
		})
	})
})

func getConfig(f *envoyhttp.HttpFilter) *envoyratelimit.RateLimit {
	cfg := f.GetConfig()
	rcfg := new(envoyratelimit.RateLimit)
	err := util.StructToMessage(cfg, rcfg)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return rcfg
}
