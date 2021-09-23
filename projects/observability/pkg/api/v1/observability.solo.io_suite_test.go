// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/k8s-utils/testutils/clusterlock"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	"github.com/solo-io/solo-kit/test/testutils"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestObservabilitySoloIo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ObservabilitySoloIo Suite")
}

var (
	lock *clusterlock.TestClusterLocker
	cfg  *rest.Config

	_ = SynchronizedAfterSuite(func() {}, func() {
		var err error
		err = os.Unsetenv(statusutils.PodNamespaceEnvName)
		Expect(err).NotTo(HaveOccurred())

		if os.Getenv("RUN_KUBE_TESTS") != "1" {
			return
		}
		ctx := context.Background()
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		clientset, err := apiexts.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "artifacts.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "authconfigs.enterprise.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "endpoints.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "gateways.gateway.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "ingresses.ingress.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "services.ingress.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "proxies.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "routeoptions.gateway.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "routetables.gateway.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "secrets.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "settings.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "upstreams.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "upstreamgroups.gloo.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "virtualhostoptions.gateway.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(ctx, "virtualservices.gateway.solo.io", metav1.DeleteOptions{})
		testutils.ErrorNotOccuredOrNotFound(err)
		Expect(lock.ReleaseLock()).NotTo(HaveOccurred())
	})

	_ = SynchronizedBeforeSuite(func() []byte {
		var err error
		err = os.Setenv(statusutils.PodNamespaceEnvName, "default")
		Expect(err).NotTo(HaveOccurred())

		if os.Getenv("RUN_KUBE_TESTS") != "1" {
			return nil
		}

		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		clientset, err := kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		lock, err = clusterlock.NewTestClusterLocker(clientset, clusterlock.Options{})
		Expect(err).NotTo(HaveOccurred())
		Expect(lock.AcquireLock()).NotTo(HaveOccurred())
		return nil
	}, func([]byte) {})
)
