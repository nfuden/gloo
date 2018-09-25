package syncer

import (
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/discovery"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/registry"
	gloosyncer "github.com/solo-io/solo-kit/projects/gloo/pkg/syncer"
)

func NewSetupSyncer() v1.SetupSyncer {
	return gloosyncer.NewSetupSyncerWithRunFunc(RunUDS)
}

func RunUDS(opts bootstrap.Opts) error {
	watchOpts := opts.WatchOpts.WithDefaults()
	watchOpts.Ctx = contextutils.WithLogger(watchOpts.Ctx, "uds")

	upstreamClient, err := v1.NewUpstreamClient(opts.Upstreams)
	if err != nil {
		return err
	}
	if err := upstreamClient.Register(); err != nil {
		return err
	}

	secretClient, err := v1.NewSecretClient(opts.Secrets)
	if err != nil {
		return err
	}
	if err := secretClient.Register(); err != nil {
		return err
	}

	emit := make(chan struct{})
	emitter := v1.NewDiscoveryEmitterWithEmit(secretClient, upstreamClient, emit)

	// jump start all the watches
	go func() {
		emit <- struct{}{}
	}()

	plugins := registry.Plugins(opts)

	var discoveryPlugins []discovery.DiscoveryPlugin
	for _, plug := range plugins {
		disc, ok := plug.(discovery.DiscoveryPlugin)
		if ok {
			discoveryPlugins = append(discoveryPlugins, disc)
		}
	}
	disc := discovery.NewUpstreamDiscovery(opts.WatchNamespaces, opts.WriteNamespace, upstreamClient, discoveryPlugins)

	sync := NewDiscoverySyncer(disc,
		discovery.Opts{}, // TODO(ilackarms)
		watchOpts.RefreshRate)
	eventLoop := v1.NewDiscoveryEventLoop(emitter, sync)

	errs := make(chan error)

	eventLoopErrs, err := eventLoop.Run(opts.WatchNamespaces, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, errs, eventLoopErrs, "event_loop.gloo")

	logger := contextutils.LoggerFrom(watchOpts.Ctx)

	for {
		select {
		case err := <-errs:
			logger.Errorf("error: %v", err)
		case <-watchOpts.Ctx.Done():
			return nil
		}
	}
}
