package v1

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/services"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("V1Cache", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1     string
		namespace2     string
		cfg            *rest.Config
		cache          Cache
		artifactClient ArtifactClient
		endpointClient EndpointClient
		proxyClient    ProxyClient
		secretClient   SecretClient
		upstreamClient UpstreamClient
		kube           kubernetes.Interface
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		err := services.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = services.SetupKubeForTest(namespace2)
		kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		Expect(err).NotTo(HaveOccurred())

		// Artifact Constructor
		kube, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		artifactClientFactory := factory.NewResourceClientFactory(&factory.KubeConfigMapClientOpts{
			Clientset: kube,
		})
		artifactClient, err = NewArtifactClient(artifactClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// Endpoint Constructor
		endpointClientFactory := factory.NewResourceClientFactory(&factory.MemoryResourceClientOpts{
			Cache: memory.NewInMemoryResourceCache(),
		})
		endpointClient, err = NewEndpointClient(endpointClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// Proxy Constructor
		proxyClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: ProxyCrd,
			Cfg: cfg,
		})
		proxyClient, err = NewProxyClient(proxyClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// Secret Constructor
		kube, err = kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())

		secretClientFactory := factory.NewResourceClientFactory(&factory.KubeSecretClientOpts{
			Clientset: kube,
		})
		secretClient, err = NewSecretClient(secretClientFactory)
		Expect(err).NotTo(HaveOccurred())

		// Upstream Constructor
		upstreamClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: UpstreamCrd,
			Cfg: cfg,
		})
		upstreamClient, err = NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())
		cache = NewCache(artifactClient, endpointClient, proxyClient, secretClient, upstreamClient)
	})
	AfterEach(func() {
		services.TeardownKube(namespace1)
		services.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := cache.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := cache.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Minute,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *Snapshot
		artifact1a, err := artifactClient.Write(NewArtifact(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact1b, err := artifactClient.Write(NewArtifact(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainartifact:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainartifact
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.Artifacts.List()).To(ContainElement(artifact1a))
		Expect(snap.Artifacts.List()).To(ContainElement(artifact1b))

		artifact2a, err := artifactClient.Write(NewArtifact(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		artifact2b, err := artifactClient.Write(NewArtifact(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Artifacts.List()).To(ContainElement(artifact1a))
			Expect(snap.Artifacts.List()).To(ContainElement(artifact1b))
			Expect(snap.Artifacts.List()).To(ContainElement(artifact2a))
			Expect(snap.Artifacts.List()).To(ContainElement(artifact2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		endpoint1a, err := endpointClient.Write(NewEndpoint(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint1b, err := endpointClient.Write(NewEndpoint(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainendpoint:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainendpoint
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.Endpoints.List()).To(ContainElement(endpoint1a))
		Expect(snap.Endpoints.List()).To(ContainElement(endpoint1b))

		endpoint2a, err := endpointClient.Write(NewEndpoint(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		endpoint2b, err := endpointClient.Write(NewEndpoint(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint1a))
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint1b))
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint2a))
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		proxy1a, err := proxyClient.Write(NewProxy(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy1b, err := proxyClient.Write(NewProxy(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainproxy:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainproxy
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.Proxies.List()).To(ContainElement(proxy1a))
		Expect(snap.Proxies.List()).To(ContainElement(proxy1b))

		proxy2a, err := proxyClient.Write(NewProxy(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		proxy2b, err := proxyClient.Write(NewProxy(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Proxies.List()).To(ContainElement(proxy1a))
			Expect(snap.Proxies.List()).To(ContainElement(proxy1b))
			Expect(snap.Proxies.List()).To(ContainElement(proxy2a))
			Expect(snap.Proxies.List()).To(ContainElement(proxy2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		secret1a, err := secretClient.Write(NewSecret(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret1b, err := secretClient.Write(NewSecret(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainsecret:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainsecret
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.Secrets.List()).To(ContainElement(secret1a))
		Expect(snap.Secrets.List()).To(ContainElement(secret1b))

		secret2a, err := secretClient.Write(NewSecret(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		secret2b, err := secretClient.Write(NewSecret(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Secrets.List()).To(ContainElement(secret1a))
			Expect(snap.Secrets.List()).To(ContainElement(secret1b))
			Expect(snap.Secrets.List()).To(ContainElement(secret2a))
			Expect(snap.Secrets.List()).To(ContainElement(secret2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		upstream1a, err := upstreamClient.Write(NewUpstream(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(NewUpstream(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainupstream:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainupstream
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.Upstreams.List()).To(ContainElement(upstream1a))
		Expect(snap.Upstreams.List()).To(ContainElement(upstream1b))

		upstream2a, err := upstreamClient.Write(NewUpstream(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(NewUpstream(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Upstreams.List()).To(ContainElement(upstream1a))
			Expect(snap.Upstreams.List()).To(ContainElement(upstream1b))
			Expect(snap.Upstreams.List()).To(ContainElement(upstream2a))
			Expect(snap.Upstreams.List()).To(ContainElement(upstream2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = artifactClient.Delete(artifact2a.Metadata.Namespace, artifact2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact2a.Metadata.Namespace, artifact2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Artifacts.List()).To(ContainElement(artifact1a))
			Expect(snap.Artifacts.List()).To(ContainElement(artifact1b))
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact2a))
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = artifactClient.Delete(artifact1a.Metadata.Namespace, artifact1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = artifactClient.Delete(artifact1b.Metadata.Namespace, artifact1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact1a))
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact1b))
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact2a))
			Expect(snap.Artifacts.List()).NotTo(ContainElement(artifact2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = endpointClient.Delete(endpoint2a.Metadata.Namespace, endpoint2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint2a.Metadata.Namespace, endpoint2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint1a))
			Expect(snap.Endpoints.List()).To(ContainElement(endpoint1b))
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint2a))
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = endpointClient.Delete(endpoint1a.Metadata.Namespace, endpoint1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = endpointClient.Delete(endpoint1b.Metadata.Namespace, endpoint1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint1a))
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint1b))
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint2a))
			Expect(snap.Endpoints.List()).NotTo(ContainElement(endpoint2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = proxyClient.Delete(proxy2a.Metadata.Namespace, proxy2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy2a.Metadata.Namespace, proxy2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Proxies.List()).To(ContainElement(proxy1a))
			Expect(snap.Proxies.List()).To(ContainElement(proxy1b))
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy2a))
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = proxyClient.Delete(proxy1a.Metadata.Namespace, proxy1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = proxyClient.Delete(proxy1b.Metadata.Namespace, proxy1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy1a))
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy1b))
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy2a))
			Expect(snap.Proxies.List()).NotTo(ContainElement(proxy2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = secretClient.Delete(secret2a.Metadata.Namespace, secret2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret2a.Metadata.Namespace, secret2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Secrets.List()).To(ContainElement(secret1a))
			Expect(snap.Secrets.List()).To(ContainElement(secret1b))
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret2a))
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = secretClient.Delete(secret1a.Metadata.Namespace, secret1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = secretClient.Delete(secret1b.Metadata.Namespace, secret1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret1a))
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret1b))
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret2a))
			Expect(snap.Secrets.List()).NotTo(ContainElement(secret2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = upstreamClient.Delete(upstream2a.Metadata.Namespace, upstream2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2a.Metadata.Namespace, upstream2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Upstreams.List()).To(ContainElement(upstream1a))
			Expect(snap.Upstreams.List()).To(ContainElement(upstream1b))
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream2a))
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = upstreamClient.Delete(upstream1a.Metadata.Namespace, upstream1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.Metadata.Namespace, upstream1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream1a))
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream1b))
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream2a))
			Expect(snap.Upstreams.List()).NotTo(ContainElement(upstream2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
	})
})
