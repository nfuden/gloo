package edge_classic_e2e

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloodefaults "github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/gloo/test/snapshot"
	"github.com/solo-io/gloo/test/snapshot/testcases"
	"github.com/solo-io/gloo/test/snapshot/utils"
	"github.com/solo-io/gloo/test/snapshot/utils/builders"
	"github.com/solo-io/go-utils/testutils"
	v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	gloov1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	"github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
	"github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/kubernetes"
	gloocore "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	gatewayPort      = int(80)
	httpbinNamespace = "httpbin"
	httpbinService   = "httpbin"
)

var _ = Describe("Gloo Edge Classic", func() {

	var (
		inputs               []client.Object
		clusterName, kubeCtx string
		resourceClientset    *kube2e.KubeResourceClientSet
		kubeClient           client.Client
		runner               snapshot.TestRunner
	)

	BeforeEach(func() {
		// Clear inputs before each run.
		inputs = []client.Object{}

		clusterName = os.Getenv("CLUSTER_NAME")
		kubeCtx = fmt.Sprintf("kind-%s", clusterName)

		// set up resource client and kubeclient
		var err error
		resourceClientset, err = kube2e.NewDefaultKubeResourceClientSet(ctx)
		Expect(err).NotTo(HaveOccurred(), "can create kube resource client set")

		kubeClient, err = utils.GetClient(kubeCtx)
		Expect(err).NotTo(HaveOccurred(), "can create client")

		runner = snapshot.TestRunner{
			Name:      "classic-apis",
			ClientSet: resourceClientset,
			Client:    kubeClient,
		}
	})

	JustAfterEach(func() {
		// Note to devs:  set NO_CLEANUP to 'all' or 'failed' to skip cleanup, for the sake of
		// debugging or otherwise examining state after a test.
		defer ctxCancel()

		if utils.ShouldSkipCleanup() {
			fmt.Printf("Not cleaning up")
			return // Exit without cleaning up
		}
		Expect(runner.Cleanup(ctx)).To(Succeed())
	})

	When("Happy Path", func() {
		BeforeEach(func() {
			vs, err := builders.NewVirtualServiceBuilder().WithName("httpbin-route").
				WithNamespace(httpbinNamespace).
				WithDomain("httpbin.example.com").
				WithRoute("httpbin-route", &v1.Route{
					Name: "httpbin-route",
					Matchers: []*matchers.Matcher{
						{
							PathSpecifier: &matchers.Matcher_Prefix{
								Prefix: "/",
							},
						},
					},
					Action: &v1.Route_RouteAction{
						RouteAction: &gloov1.RouteAction{
							Destination: &gloov1.RouteAction_Single{
								Single: &gloov1.Destination{
									DestinationType: &gloov1.Destination_Upstream{
										Upstream: &gloocore.ResourceRef{
											Name:      "httpbin-htppbin-8000",
											Namespace: gloodefaults.GlooSystem,
										},
									},
								},
							},
						},
					},
				}).Build()
			Expect(err).NotTo(HaveOccurred())

			inputs = []client.Object{
				builders.NewUpstreamBuilder().
					WithName("httpbin-htppbin-8000").
					WithNamespace(gloodefaults.GlooSystem).
					WithDiscoveryMetadata(&gloov1.DiscoveryMetadata{
						Labels: map[string]string{
							"app":     httpbinService,
							"service": httpbinService,
						},
					}).
					WithKubeUpstream(&gloov1.UpstreamSpec_Kube{
						Kube: &kubernetes.UpstreamSpec{
							Selector: map[string]string{
								"app": httpbinService,
							},
							ServiceNamespace: httpbinNamespace,
							ServiceName:      httpbinService,
							ServicePort:      uint32(8000),
						},
					}).
					Build(),
				vs,
			}
		})

		It("Send request through ingress", func() {
			testcases.TestGatewayIngress(
				ctx,
				runner,
				&snapshot.TestEnv{
					GatewayName:      defaults.GatewayProxyName,
					GatewayNamespace: gloodefaults.GlooSystem,
					GatewayPort:      gatewayPort,
					ClusterName:      clusterName,
					ClusterContext:   kubeCtx,
				},
				inputs,
				func() {
					err := testutils.WaitPodsRunning(ctx, time.Second, gloodefaults.GlooSystem, fmt.Sprintf("gloo=%s", defaults.GatewayProxyName))
					Expect(err).NotTo(HaveOccurred())
				},
			)
		})
	})
})
