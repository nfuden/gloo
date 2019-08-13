// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"context"
	"net"

	"github.com/solo-io/go-utils/envutils"
	"github.com/solo-io/solo-projects/pkg/license"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/helpers/rawgetter"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/kube"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/settings"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/artifactsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/configsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/gatewaysvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/proxysvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/secretsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc/mutation"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/converter"
	mutation2 "github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mutation"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/selection"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/setup"
)

// Injectors from wire.go:

func InitializeServer(ctx context.Context, listener net.Listener) (*GlooGrpcService, error) {
	v1Settings := setup.MustSettings(ctx)
	clientSet, err := setup.NewClientSet(ctx, v1Settings)
	if err != nil {
		return nil, err
	}
	upstreamClient := setup.NewUpstreamClient(clientSet)
	settingsClient := setup.NewSettingsClient(clientSet)
	valuesClient := settings.NewSettingsValuesClient(ctx, settingsClient)
	mutator := mutation.NewMutator(upstreamClient)
	factory := mutation.NewFactory()
	upstreamApiServer := upstreamsvc.NewUpstreamGrpcService(ctx, upstreamClient, valuesClient, mutator, factory)
	artifactClient := setup.NewArtifactClient(clientSet)
	artifactApiServer := artifactsvc.NewArtifactGrpcService(ctx, artifactClient)
	client := license.NewClient(ctx)
	coreV1Interface := setup.NewCoreV1Interface(clientSet)
	namespaceClient := kube.NewNamespaceClient(coreV1Interface)
	oAuthEndpoint := setup.NewOAuthEndpoint()
	buildVersion := setup.GetBuildVersion()
	string2 := envutils.MustGetPodNamespace(ctx)
	configApiServer := configsvc.NewConfigGrpcService(ctx, settingsClient, client, namespaceClient, oAuthEndpoint, buildVersion, string2)
	secretClient := setup.NewSecretClient(clientSet)
	secretApiServer := secretsvc.NewSecretGrpcService(ctx, secretClient)
	virtualServiceClient := setup.NewVirtualServiceClient(clientSet)
	mutationMutator := mutation2.NewMutator(ctx, virtualServiceClient)
	mutationFactory := mutation2.NewMutationFactory()
	rawGetter := rawgetter.NewKubeYamlRawGetter()
	virtualServiceDetailsConverter := converter.NewVirtualServiceDetailsConverter(rawGetter)
	virtualServiceSelector := selection.NewVirtualServiceSelector(virtualServiceClient, namespaceClient, string2)
	virtualServiceApiServer := virtualservicesvc.NewVirtualServiceGrpcService(ctx, string2, virtualServiceClient, valuesClient, mutationMutator, mutationFactory, virtualServiceDetailsConverter, virtualServiceSelector)
	gatewayClient := setup.NewGatewayClient(clientSet)
	gatewayApiServer := gatewaysvc.NewGatewayGrpcService(ctx, gatewayClient, rawGetter)
	proxyClient := setup.NewProxyClient(clientSet)
	proxyApiServer := proxysvc.NewProxyGrpcService(ctx, proxyClient, rawGetter)
	glooGrpcService := NewGlooGrpcService(listener, upstreamApiServer, artifactApiServer, configApiServer, secretApiServer, virtualServiceApiServer, gatewayApiServer, proxyApiServer)
	return glooGrpcService, nil
}
