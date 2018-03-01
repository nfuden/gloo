package controller

import (
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	kubeplugin "github.com/solo-io/gloo-plugins/kubernetes"
	"github.com/solo-io/gloo-storage"
	"github.com/solo-io/gloo-storage/crd"
	. "github.com/solo-io/gloo-testing/helpers"
)

var _ = Describe("Kube Service Discovery", func() {
	var (
		masterUrl, kubeconfigPath string
		mkb                       *MinikubeInstance
		namespace                 string
	)
	BeforeEach(func() {
		namespace = RandString(8)
		mkb = NewMinikube(false, namespace)
		err := mkb.Setup()
		Must(err)
		kubeconfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		masterUrl, err = mkb.Addr()
		Must(err)
	})
	AfterEach(func() {
		mkb.Teardown()
	})
	Describe("service discovery", func() {
		var (
			discovery  *ServiceController
			kubeClient kubernetes.Interface
			glooClient storage.Interface
		)
		BeforeEach(func() {
			cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
			Must(err)

			glooClient, err = crd.NewStorage(cfg, namespace, time.Minute)
			Must(err)

			discovery, err = NewServiceController(cfg, glooClient, time.Minute)
			Must(err)

			go discovery.Run(make(chan struct{}))

			kubeClient, err = kubernetes.NewForConfig(cfg)
			Must(err)
		})
		Context("a service exists", func() {
			var (
				serviceName string
				service     *kubev1.Service
			)
			BeforeEach(func() {
				serviceName = "somethingsomethingsomething"
				service = &kubev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      serviceName,
						Namespace: namespace,
					},
					Spec: kubev1.ServiceSpec{
						Ports: []kubev1.ServicePort{
							{
								Name: "foo",
								Port: 8080,
							},
						},
					},
				}
				_, err := kubeClient.CoreV1().Services(namespace).Create(service)
				Expect(err).NotTo(HaveOccurred())
			})
			AfterEach(func() {
				err := kubeClient.CoreV1().Services(namespace).Delete(serviceName, nil)
				Expect(err).NotTo(HaveOccurred())
			})
			It("does not return an error", func() {
				select {
				case <-time.After(time.Second):
					// passed without error
				case err := <-discovery.Error():
					Expect(err).NotTo(HaveOccurred())
					Fail("err passed, but was nil")
				}
			})
			It("creates an upstream for each port from the service definition", func() {
				time.Sleep(time.Second * 3)
				createdUpstreams, err := glooClient.V1().Upstreams().List()
				// need to clear metadata for test
				for _, us := range createdUpstreams {
					us.Metadata = nil
				}
				Expect(err).NotTo(HaveOccurred())
				for _, port := range service.Spec.Ports {
					expectedUpstream := &v1.Upstream{
						Name: upstreamName(namespace, serviceName, port.Port),
						Type: kubeplugin.UpstreamTypeKube,
						Spec: kubeplugin.EncodeUpstreamSpec(kubeplugin.UpstreamSpec{
							ServiceNamespace: namespace,
							ServiceName:      serviceName,
							ServicePort:      port.TargetPort.IntVal,
						}),
					}
					Expect(createdUpstreams).To(ContainElement(expectedUpstream))
				}
			})
		})
	})
})
