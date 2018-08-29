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
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/services"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("V1Cache", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1        string
		namespace2        string
		cfg               *rest.Config
		cache             Cache
		resolverMapClient ResolverMapClient
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

		// ResolverMap Constructor
		resolverMapClientFactory := factory.NewResourceClientFactory(&factory.KubeResourceClientOpts{
			Crd: ResolverMapCrd,
			Cfg: cfg,
		})
		resolverMapClient, err = NewResolverMapClient(resolverMapClientFactory)
		Expect(err).NotTo(HaveOccurred())
		cache = NewCache(resolverMapClient)
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
		resolverMap1a, err := resolverMapClient.Write(NewResolverMap(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		resolverMap1b, err := resolverMapClient.Write(NewResolverMap(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

	drainresolverMap:
		for {
			select {
			case snap = <-snapshots:
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Millisecond * 500):
				break drainresolverMap
			case <-time.After(time.Second):
				Fail("expected snapshot before 1 second")
			}
		}
		Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1a))
		Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1b))

		resolverMap2a, err := resolverMapClient.Write(NewResolverMap(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		resolverMap2b, err := resolverMapClient.Write(NewResolverMap(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1a))
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1b))
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap2a))
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
		err = resolverMapClient.Delete(resolverMap2a.Metadata.Namespace, resolverMap2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = resolverMapClient.Delete(resolverMap2a.Metadata.Namespace, resolverMap2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1a))
			Expect(snap.ResolverMaps.List()).To(ContainElement(resolverMap1b))
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap2a))
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}

		err = resolverMapClient.Delete(resolverMap1a.Metadata.Namespace, resolverMap1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = resolverMapClient.Delete(resolverMap1b.Metadata.Namespace, resolverMap1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		select {
		case snap := <-snapshots:
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap1a))
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap1b))
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap2a))
			Expect(snap.ResolverMaps.List()).NotTo(ContainElement(resolverMap2b))
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Second * 3):
			Fail("expected snapshot before 1 second")
		}
	})
})
