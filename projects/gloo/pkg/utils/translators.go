package utils

import (
	"fmt"
	"strings"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	// label key applied to Proxies generated by the Gloo Edge translator
	ProxyTypeKey = "proxy_type"
	// label value applied to Proxies generated by the Gloo Edge translator
	GlooEdgeTranslatorValue = "gloo-gateway"
	// label value applied to Proxies generated by the Gloo Gateway translator
	GlooGatewayTranslatorValue = "gloo-kube-gateway-api"

	// Gloo Gateway proxies can live in different namespaces from writeNamespace
	NamespaceLabel = "gateway.solo.io/namespace"
)

func GetTranslatorSelectorExpression(translators ...string) string {
	return fmt.Sprintf("%s in (%s)", ProxyTypeKey, strings.Join(translators, ", "))
}

func GetTranslatorValue(meta *core.Metadata) string {
	return meta.GetLabels()[ProxyTypeKey]
}
