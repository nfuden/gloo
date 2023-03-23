// Code generated by skv2. DO NOT EDIT.

package federated_gloo_resource_handler

import (
	"context"

	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"gopkg.in/yaml.v1"

	"github.com/solo-io/go-utils/contextutils"
	rpc_fed_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	gloo_solo_io_v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gloo.solo.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewFederatedGlooResourceHandler(
	glooFedClient gloo_solo_io_v1.Clientset,

) rpc_fed_v1.FederatedGlooResourceApiServer {
	return &glooFedResourceHandler{
		glooFedClient: glooFedClient,
	}
}

type glooFedResourceHandler struct {
	glooFedClient gloo_solo_io_v1.Clientset
}

func (k *glooFedResourceHandler) ListFederatedUpstreams(ctx context.Context, request *rpc_fed_v1.ListFederatedUpstreamsRequest) (*rpc_fed_v1.ListFederatedUpstreamsResponse, error) {
	var rpcFederatedUpstreams []*rpc_fed_v1.FederatedUpstream
	list, err := k.glooFedClient.FederatedUpstreams().ListFederatedUpstream(ctx)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to list federatedUpstream")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err))
		return nil, wrapped
	}
	for i, _ := range list.Items {
		rpcFederatedUpstreams = append(rpcFederatedUpstreams, BuildRpcFederatedUpstream(&list.Items[i]))
	}
	return &rpc_fed_v1.ListFederatedUpstreamsResponse{
		FederatedUpstreams: rpcFederatedUpstreams,
	}, nil
}

func BuildRpcFederatedUpstream(federatedUpstream *gloo_solo_io_v1.FederatedUpstream) *rpc_fed_v1.FederatedUpstream {
	return &rpc_fed_v1.FederatedUpstream{
		Metadata: apiserverutils.ToMetadata(federatedUpstream.ObjectMeta),
		Spec:     &federatedUpstream.Spec,
		Status:   &federatedUpstream.Status,
	}
}

func (k *glooFedResourceHandler) GetFederatedUpstreamYaml(ctx context.Context, request *rpc_fed_v1.GetFederatedUpstreamYamlRequest) (*rpc_fed_v1.GetFederatedUpstreamYamlResponse, error) {
	federatedUpstream, err := k.glooFedClient.FederatedUpstreams().GetFederatedUpstream(ctx, client.ObjectKey{
		Namespace: request.GetFederatedUpstreamRef().GetNamespace(),
		Name:      request.GetFederatedUpstreamRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get FederatedUpstream")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(federatedUpstream)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_fed_v1.GetFederatedUpstreamYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *glooFedResourceHandler) ListFederatedUpstreamGroups(ctx context.Context, request *rpc_fed_v1.ListFederatedUpstreamGroupsRequest) (*rpc_fed_v1.ListFederatedUpstreamGroupsResponse, error) {
	var rpcFederatedUpstreamGroups []*rpc_fed_v1.FederatedUpstreamGroup
	list, err := k.glooFedClient.FederatedUpstreamGroups().ListFederatedUpstreamGroup(ctx)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to list federatedUpstreamGroup")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err))
		return nil, wrapped
	}
	for i, _ := range list.Items {
		rpcFederatedUpstreamGroups = append(rpcFederatedUpstreamGroups, BuildRpcFederatedUpstreamGroup(&list.Items[i]))
	}
	return &rpc_fed_v1.ListFederatedUpstreamGroupsResponse{
		FederatedUpstreamGroups: rpcFederatedUpstreamGroups,
	}, nil
}

func BuildRpcFederatedUpstreamGroup(federatedUpstreamGroup *gloo_solo_io_v1.FederatedUpstreamGroup) *rpc_fed_v1.FederatedUpstreamGroup {
	return &rpc_fed_v1.FederatedUpstreamGroup{
		Metadata: apiserverutils.ToMetadata(federatedUpstreamGroup.ObjectMeta),
		Spec:     &federatedUpstreamGroup.Spec,
		Status:   &federatedUpstreamGroup.Status,
	}
}

func (k *glooFedResourceHandler) GetFederatedUpstreamGroupYaml(ctx context.Context, request *rpc_fed_v1.GetFederatedUpstreamGroupYamlRequest) (*rpc_fed_v1.GetFederatedUpstreamGroupYamlResponse, error) {
	federatedUpstreamGroup, err := k.glooFedClient.FederatedUpstreamGroups().GetFederatedUpstreamGroup(ctx, client.ObjectKey{
		Namespace: request.GetFederatedUpstreamGroupRef().GetNamespace(),
		Name:      request.GetFederatedUpstreamGroupRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get FederatedUpstreamGroup")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(federatedUpstreamGroup)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_fed_v1.GetFederatedUpstreamGroupYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *glooFedResourceHandler) ListFederatedSettings(ctx context.Context, request *rpc_fed_v1.ListFederatedSettingsRequest) (*rpc_fed_v1.ListFederatedSettingsResponse, error) {
	var rpcFederatedSettings []*rpc_fed_v1.FederatedSettings
	list, err := k.glooFedClient.FederatedSettings().ListFederatedSettings(ctx)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to list federatedSettings")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err))
		return nil, wrapped
	}
	for i, _ := range list.Items {
		rpcFederatedSettings = append(rpcFederatedSettings, BuildRpcFederatedSettings(&list.Items[i]))
	}
	return &rpc_fed_v1.ListFederatedSettingsResponse{
		FederatedSettings: rpcFederatedSettings,
	}, nil
}

func BuildRpcFederatedSettings(federatedSettings *gloo_solo_io_v1.FederatedSettings) *rpc_fed_v1.FederatedSettings {
	return &rpc_fed_v1.FederatedSettings{
		Metadata: apiserverutils.ToMetadata(federatedSettings.ObjectMeta),
		Spec:     &federatedSettings.Spec,
		Status:   &federatedSettings.Status,
	}
}

func (k *glooFedResourceHandler) GetFederatedSettingsYaml(ctx context.Context, request *rpc_fed_v1.GetFederatedSettingsYamlRequest) (*rpc_fed_v1.GetFederatedSettingsYamlResponse, error) {
	federatedSettings, err := k.glooFedClient.FederatedSettings().GetFederatedSettings(ctx, client.ObjectKey{
		Namespace: request.GetFederatedSettingsRef().GetNamespace(),
		Name:      request.GetFederatedSettingsRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get FederatedSettings")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(federatedSettings)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_fed_v1.GetFederatedSettingsYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}