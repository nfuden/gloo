package basic_test

import (
	"context"
	"time"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	k8s_core_v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Discovery e2e", func() {

	var (
		restCfg *rest.Config
	)

	BeforeEach(func() {
		restCfg, err = config.GetConfigWithContext("")
		Expect(err).NotTo(HaveOccurred())
	})

	It("works for discovering the remote gloo instance installed by the test harness", func() {
		clientset, err := v1.NewClientsetFromConfig(restCfg)
		Expect(err).NotTo(HaveOccurred())

		var instance *v1.GlooInstance
		Eventually(func() *v1.GlooInstance {
			instance, _ = clientset.GlooInstances().
				GetGlooInstance(
					context.TODO(),
					types.NamespacedName{
						Name:      remoteClusterConfig.KubeContext + "-gloo-system",
						Namespace: "gloo-system",
					})

			return instance
		}, 5*time.Second).ShouldNot(BeNil())

		Expect(instance.Spec.GetCluster()).To(BeEquivalentTo(remoteClusterConfig.KubeContext))
		Expect(instance.Spec.GetControlPlane().GetNamespace()).To(BeEquivalentTo("gloo-system"))
		Expect(instance.Spec.GetControlPlane().GetWatchedNamespaces()).To(BeEquivalentTo([]string{"gloo-system", "default"}))
		Expect(instance.Spec.IsEnterprise).To(BeEquivalentTo(true))
		Expect(instance.Spec.GetProxies()).To(HaveLen(1))
		Expect(instance.Spec.GetProxies()[0].GetNamespace()).To(BeEquivalentTo("gloo-system"))
		Expect(instance.Spec.GetProxies()[0].GetReadyReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetAvailableReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetWasmEnabled()).To(BeFalse())
		Expect(instance.Spec.GetProxies()[0].GetIngressEndpoints()).To(HaveLen(1))
		// check gateway-proxy, gateway-proxy-ssl and gateway-proxy-failover gateways are created
		Expect(instance.Spec.GetCheck().GetGateways().GetTotal()).To(BeEquivalentTo(3))

		coreRemoteClientset, err := k8s_core_v1.NewClientsetFromConfig(remoteClusterConfig.RestConfig)
		Expect(err).NotTo(HaveOccurred())

		ingressEndpoints := instance.Spec.GetProxies()[0].GetIngressEndpoints()[0]
		Expect(ingressEndpoints.GetServiceName()).To(Equal("gateway-proxy"))
		Expect(ingressEndpoints.GetPorts()).To(HaveLen(3))
		gpSvc, err := coreRemoteClientset.Services().GetService(context.TODO(), client.ObjectKey{
			Namespace: "gloo-system",
			Name:      ingressEndpoints.GetServiceName(),
		})
		Expect(err).NotTo(HaveOccurred())
		for _, port := range ingressEndpoints.GetPorts() {
			var discoveredPort bool
			for _, svcPort := range gpSvc.Spec.Ports {
				if svcPort.Name == port.GetName() && uint32(svcPort.NodePort) == port.GetPort() {
					discoveredPort = true
					break
				}
			}
			Expect(discoveredPort).To(BeTrue(), "found the gateway-proxy port entry in the port list")
		}
	})

	It("works for discovering the gloo-ee instance installed on the local cluster by the test harness", func() {
		clientset, err := v1.NewClientsetFromConfig(restCfg)
		Expect(err).NotTo(HaveOccurred())
		var instance *v1.GlooInstance
		Eventually(func() *v1.GlooInstance {
			instance, _ = clientset.GlooInstances().
				GetGlooInstance(
					context.TODO(),
					types.NamespacedName{
						Name:      managementClusterConfig.KubeContext + "-gloo-system",
						Namespace: "gloo-system",
					})
			return instance
		}, 5*time.Second).ShouldNot(BeNil())
		Expect(instance.Spec.GetCluster()).To(BeEquivalentTo(managementClusterConfig.KubeContext))
		Expect(instance.Spec.GetControlPlane().GetNamespace()).To(BeEquivalentTo("gloo-system"))
		Expect(instance.Spec.GetControlPlane().GetWatchedNamespaces()).To(HaveLen(0))
		Expect(instance.Spec.IsEnterprise).To(BeEquivalentTo(true))
		Expect(instance.Spec.GetProxies()).To(HaveLen(1))
		Expect(instance.Spec.GetProxies()[0].GetNamespace()).To(BeEquivalentTo("gloo-system"))
		Expect(instance.Spec.GetProxies()[0].GetReadyReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetAvailableReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetReplicas()).To(BeEquivalentTo(1))
		Expect(instance.Spec.GetProxies()[0].GetWasmEnabled()).To(BeFalse())
		Expect(instance.Spec.GetProxies()[0].GetIngressEndpoints()).To(HaveLen(1))
		Expect(instance.Spec.GetCheck().GetGateways().GetTotal()).To(BeEquivalentTo(2))
	})

})