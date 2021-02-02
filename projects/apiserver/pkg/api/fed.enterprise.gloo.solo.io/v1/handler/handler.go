// Code generated by skv2. DO NOT EDIT.

package federated_enterprise_gloo_resource_handler

import (
	"context"

	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"gopkg.in/yaml.v1"

	"github.com/solo-io/go-utils/contextutils"
	rpc_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.enterprise.gloo.solo.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewFederatedEnterpriseGlooResourceHandler(
	enterprise_glooFedClient enterprise_gloo_solo_io_v1.Clientset,

) rpc_v1.FederatedEnterpriseGlooResourceApiServer {
	return &enterprise_glooFedResourceHandler{
		enterprise_glooFedClient: enterprise_glooFedClient,
	}
}

type enterprise_glooFedResourceHandler struct {
	enterprise_glooFedClient enterprise_gloo_solo_io_v1.Clientset
}

func (k *enterprise_glooFedResourceHandler) ListFederatedAuthConfigs(ctx context.Context, request *rpc_v1.ListFederatedAuthConfigsRequest) (*rpc_v1.ListFederatedAuthConfigsResponse, error) {
	var rpcFederatedAuthConfigs []*rpc_v1.FederatedAuthConfig
	list, err := k.enterprise_glooFedClient.FederatedAuthConfigs().ListFederatedAuthConfig(ctx)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to list federatedAuthConfig")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err))
		return nil, wrapped
	}
	for i, _ := range list.Items {
		rpcFederatedAuthConfigs = append(rpcFederatedAuthConfigs, BuildRpcFederatedAuthConfig(&list.Items[i]))
	}
	return &rpc_v1.ListFederatedAuthConfigsResponse{
		FederatedAuthConfigs: rpcFederatedAuthConfigs,
	}, nil
}

func BuildRpcFederatedAuthConfig(federatedAuthConfig *enterprise_gloo_solo_io_v1.FederatedAuthConfig) *rpc_v1.FederatedAuthConfig {
	return &rpc_v1.FederatedAuthConfig{
		Metadata: apiserverutils.ToMetadata(federatedAuthConfig.ObjectMeta),
		Spec:     &federatedAuthConfig.Spec,
		Status:   &federatedAuthConfig.Status,
	}
}

func (k *enterprise_glooFedResourceHandler) GetFederatedAuthConfigYaml(ctx context.Context, request *rpc_v1.GetFederatedAuthConfigYamlRequest) (*rpc_v1.GetFederatedAuthConfigYamlResponse, error) {
	federatedAuthConfig, err := k.enterprise_glooFedClient.FederatedAuthConfigs().GetFederatedAuthConfig(ctx, client.ObjectKey{
		Namespace: request.GetFederatedAuthConfigRef().GetNamespace(),
		Name:      request.GetFederatedAuthConfigRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get FederatedAuthConfig")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(federatedAuthConfig)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_v1.GetFederatedAuthConfigYamlResponse{
		YamlData: &rpc_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}
