// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	"github.com/solo-io/solo-kit/test/helpers"
	apiext "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		ctx            context.Context
		namespace1     string
		namespace2     string
		name1, name2   = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg            *rest.Config
		clientset      *apiext.Clientset
		kube           kubernetes.Interface
		emitter        DashboardsEmitter
		upstreamClient gloo_solo_io.UpstreamClient
	)

	BeforeEach(func() {
		err := os.Setenv(statusutils.PodNamespaceEnvName, "default")
		Expect(err).NotTo(HaveOccurred())

		ctx = context.Background()
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		kube = helpers.MustKubeClient()
		err = kubeutils.CreateNamespacesInParallel(ctx, kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		clientset, err = apiext.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		// Upstream Constructor
		upstreamClientFactory := &factory.KubeResourceClientFactory{
			Crd:         gloo_solo_io.UpstreamCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		err = helpers.AddAndRegisterCrd(ctx, gloo_solo_io.UpstreamCrd, clientset)
		Expect(err).NotTo(HaveOccurred())

		upstreamClient, err = gloo_solo_io.NewUpstreamClient(ctx, upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewDashboardsEmitter(upstreamClient)
	})
	AfterEach(func() {
		err := os.Unsetenv(statusutils.PodNamespaceEnvName)
		Expect(err).NotTo(HaveOccurred())

		err = kubeutils.DeleteNamespacesInParallelBlocking(ctx, kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
	})

	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *DashboardsSnapshot

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams gloo_solo_io.UpstreamList, unexpectUpstreams gloo_solo_io.UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, gloo_solo_io.UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})
	})

	It("tracks snapshots on changes to any resource using AllNamespace", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{""}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *DashboardsSnapshot

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams gloo_solo_io.UpstreamList, unexpectUpstreams gloo_solo_io.UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, gloo_solo_io.UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})
	})
})
