package template

import (
	"testing"

	kurlv1beta1 "github.com/replicatedhq/kurl/kurlkinds/pkg/apis/cluster/v1beta1"
	"github.com/stretchr/testify/require"
	"go.undefinedlabs.com/scopeagent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBoolPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = true

	outcome := ctx.kurlBool("test")
	req.Equal(outcome, true)
}

func TestBoolNotPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = true

	outcome := ctx.kurlBool("wrong")
	req.NotEqual(outcome, true)
}

func TestBoolInvalidType(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = 6

	outcome := ctx.kurlBool("test")
	req.NotEqual(outcome, true)
}

func TestStringPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = "value"

	outcome := ctx.kurlString("test")
	req.Equal(outcome, "value")
}

func TestStringNotPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = "value"

	outcome := ctx.kurlString("wrong")
	req.NotEqual(outcome, "value")
}

func TestStringInvalidType(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = 6

	outcome := ctx.kurlString("test")
	req.Equal(outcome, "")
}

func TestIntPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = 42

	outcome := ctx.kurlInt("test")
	req.Equal(outcome, 42)
}

func TestIntNotPresent(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = 42

	outcome := ctx.kurlInt("wrong")
	req.NotEqual(outcome, 42)
}

func TestIntInvalidType(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	ctx.KurlValues["test"] = false

	outcome := ctx.kurlInt("test")
	req.Equal(outcome, 0)
}

func TestParseInstallerProperly(t *testing.T) {
	scopetest := scopeagent.StartTest(t)
	defer scopetest.End()
	req := require.New(t)

	ctx := &KurlCtx{
		KurlValues: make(map[string]interface{}),
	}

	kurlInstaller := &kurlv1beta1.Installer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: kurlv1beta1.InstallerSpec{
			Kubernetes: kurlv1beta1.Kubernetes{
				MasterAddress:    "1.1.1.1",
				ServiceCIDR:      "/24",
				ServiceCidrRange: "1.1.1.1",
				Version:          "latest",
			},
			Fluentd: kurlv1beta1.Fluentd{
				FullEFKStack: true,
				Version:      "latest",
			},
			Kotsadm: kurlv1beta1.Kotsadm{
				ApplicationNamespace: "namelike",
				ApplicationSlug:      "sluglike",
				Hostname:             "104.24.13.4",
				UiBindPort:           24,
				Version:              "latest",
			},
		},
	}

	ctx.AddValuesToKurlContext(kurlInstaller)

	// The kurlString, kurlInt, and kurlBool methods accept a
	// yamlPath delimeted by a '.' to reach a resource

	kubernetesVersion := ctx.kurlString("Kubernetes.Version")
	req.Equal(kubernetesVersion, "latest")

	kotsadmUiBindPort := ctx.kurlInt("Kotsadm.UiBindPort")
	req.Equal(kotsadmUiBindPort, 24)

	fluentdFullEFKStack := ctx.kurlBool("Fluentd.FullEFKStack")
	req.Equal(fluentdFullEFKStack, true)
}
