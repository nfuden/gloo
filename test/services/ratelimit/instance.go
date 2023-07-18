package ratelimit

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/gomega"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/rate-limiter/pkg/server"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-projects/projects/rate-limit/pkg/runner"
	"github.com/solo-io/solo-projects/projects/rate-limit/pkg/xds"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Instance struct {
	address string

	serverSettings server.Settings
}

func (i *Instance) UpdateServerSettings(mutator func(*server.Settings)) {
	mutator(&i.serverSettings)
}

func (i *Instance) RunWithXds(ctx context.Context, xdsPort uint32) {
	runner.Run(ctx, i.serverSettings, xds.Settings{
		GlooAddress: fmt.Sprintf("%s:%d", i.Address(), xdsPort),
	})
}

func (i *Instance) Address() string {
	return i.address
}

func (i *Instance) Port() uint32 {
	return uint32(i.serverSettings.RateLimitPort)
}

func (i *Instance) Url() string {
	return fmt.Sprintf("%s:%d", i.Address(), i.Port())
}

func (i *Instance) GetHealthCheckResponse(opts ...grpc.CallOption) (*healthpb.HealthCheckResponse, error) {
	conn, err := grpc.Dial(i.Url(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	return healthpb.NewHealthClient(conn).Check(context.Background(), &healthpb.HealthCheckRequest{
		Service: i.serverSettings.GrpcServiceName,
	}, opts...)
}

func (i *Instance) IsHealthy() bool {
	response, err := i.GetHealthCheckResponse()
	if err != nil {
		return false
	}
	if response.Status != healthpb.HealthCheckResponse_SERVING {
		return false
	}
	return true
}

func (i *Instance) EventuallyIsHealthy() {
	EventuallyWithOffset(1, func(g Gomega) {
		g.Expect(i.IsHealthy()).To(BeTrue())
	}, "5s", ".1s").Should(Succeed())
}

func (i *Instance) GetServerUpstream() *gloov1.Upstream {
	return &gloov1.Upstream{
		Metadata: &core.Metadata{
			Name:      "rl-server",
			Namespace: defaults.GlooSystem,
		},
		UseHttp2: &wrappers.BoolValue{Value: true},
		UpstreamType: &gloov1.Upstream_Static{
			Static: &gloov1static.UpstreamSpec{
				Hosts: []*gloov1static.Host{{
					Addr: i.Address(),
					Port: i.Port(),
				}},
			},
		},
	}
}