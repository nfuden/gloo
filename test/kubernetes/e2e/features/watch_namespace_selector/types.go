package watch_namespace_selector

import (
	"path/filepath"

	"github.com/solo-io/gloo/test/kubernetes/e2e/tests/base"
	"github.com/solo-io/skv2/codegen/util"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	setupManifest       = filepath.Join(util.MustGetThisDir(), "testdata", "setup.yaml")
	installNSVSManifest = filepath.Join(util.MustGetThisDir(), "testdata", "vs-install-ns.yaml")

	matchLabelsSetup = filepath.Join(util.MustGetThisDir(), "testdata", "match-labels.yaml")

	unlabeledRandomNamespaceManifest = filepath.Join(util.MustGetThisDir(), "testdata", "random-ns-unlabeled.yaml")
	labeledRandomNamespaceManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "random-ns-labeled.yaml")
	randomVSManifest                 = filepath.Join(util.MustGetThisDir(), "testdata", "vs-random.yaml")

	randomNSVS = &gatewayv1.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "vs-random",
			Namespace: "random",
		},
	}

	setupSuite = base.SimpleTestCase{
		Manifests: []string{setupManifest},
	}

	testCases = map[string]*base.TestCase{
		"TestMatchLabels": {
			SimpleTestCase: base.SimpleTestCase{
				UpgradeValues: matchLabelsSetup,
				Manifests:     []string{unlabeledRandomNamespaceManifest, randomVSManifest},
				Resources:     []client.Object{randomNSVS},
			},
		},
	}
)
