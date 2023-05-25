// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mDiscoveryResourcesIn
	mDiscoverySnapshotIn = stats.Int64("discovery.gloo.solo.io/emitter/snap_in", "Deprecated. Use discovery.gloo.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mDiscoveryResourcesIn    = stats.Int64("discovery.gloo.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mDiscoverySnapshotOut    = stats.Int64("discovery.gloo.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mDiscoverySnapshotMissed = stats.Int64("discovery.gloo.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see discoveryResourcesInView
	discoverysnapshotInView = &view.View{
		Name:        "discovery.gloo.solo.io/emitter/snap_in",
		Measure:     mDiscoverySnapshotIn,
		Description: "Deprecated. Use discovery.gloo.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	discoveryResourcesInView = &view.View{
		Name:        "discovery.gloo.solo.io/emitter/resources_in",
		Measure:     mDiscoveryResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	discoverysnapshotOutView = &view.View{
		Name:        "discovery.gloo.solo.io/emitter/snap_out",
		Measure:     mDiscoverySnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	discoverysnapshotMissedView = &view.View{
		Name:        "discovery.gloo.solo.io/emitter/snap_missed",
		Measure:     mDiscoverySnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		discoverysnapshotInView,
		discoverysnapshotOutView,
		discoverysnapshotMissedView,
		discoveryResourcesInView,
	)
}

type DiscoverySnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *DiscoverySnapshot, <-chan error, error)
}

type DiscoveryEmitter interface {
	DiscoverySnapshotEmitter
	Register() error
	Upstream() UpstreamClient
	KubeNamespace() github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceClient
	Secret() SecretClient
}

func NewDiscoveryEmitter(upstreamClient UpstreamClient, kubeNamespaceClient github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceClient, secretClient SecretClient) DiscoveryEmitter {
	return NewDiscoveryEmitterWithEmit(upstreamClient, kubeNamespaceClient, secretClient, make(chan struct{}))
}

func NewDiscoveryEmitterWithEmit(upstreamClient UpstreamClient, kubeNamespaceClient github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceClient, secretClient SecretClient, emit <-chan struct{}) DiscoveryEmitter {
	return &discoveryEmitter{
		upstream:      upstreamClient,
		kubeNamespace: kubeNamespaceClient,
		secret:        secretClient,
		forceEmit:     emit,
	}
}

type discoveryEmitter struct {
	forceEmit     <-chan struct{}
	upstream      UpstreamClient
	kubeNamespace github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceClient
	secret        SecretClient
}

func (c *discoveryEmitter) Register() error {
	if err := c.upstream.Register(); err != nil {
		return err
	}
	if err := c.kubeNamespace.Register(); err != nil {
		return err
	}
	if err := c.secret.Register(); err != nil {
		return err
	}
	return nil
}

func (c *discoveryEmitter) Upstream() UpstreamClient {
	return c.upstream
}

func (c *discoveryEmitter) KubeNamespace() github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceClient {
	return c.kubeNamespace
}

func (c *discoveryEmitter) Secret() SecretClient {
	return c.secret
}

func (c *discoveryEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *DiscoverySnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	var initialUpstreamList UpstreamList
	/* Create channel for KubeNamespace */
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)

	var initialSecretList SecretList

	currentSnapshot := DiscoverySnapshot{}
	upstreamsByNamespace := make(map[string]UpstreamList)
	secretsByNamespace := make(map[string]SecretList)

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Upstream */
		{
			upstreams, err := c.upstream.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Upstream list")
			}
			initialUpstreamList = append(initialUpstreamList, upstreams...)
			upstreamsByNamespace[namespace] = upstreams
		}
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)
		/* Setup namespaced watch for Secret */
		{
			secrets, err := c.secret.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Secret list")
			}
			initialSecretList = append(initialSecretList, secrets...)
			secretsByNamespace[namespace] = secrets
		}
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case upstreamList, ok := <-upstreamNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				case secretList, ok := <-secretNamespacesChan:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Upstreams */
	currentSnapshot.Upstreams = initialUpstreamList.Sort()
	/* Setup cluster-wide watch for KubeNamespace */
	var err error
	currentSnapshot.Kubenamespaces, err = c.kubeNamespace.List(clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "initial KubeNamespace list")
	}
	kubeNamespaceChan, kubeNamespaceErrs, err := c.kubeNamespace.Watch(opts)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "starting KubeNamespace watch")
	}
	done.Add(1)
	go func() {
		defer done.Done()
		errutils.AggregateErrs(ctx, errs, kubeNamespaceErrs, "kubenamespaces")
	}()
	/* Initialize snapshot for Secrets */
	currentSnapshot.Secrets = initialSecretList.Sort()

	snapshots := make(chan *DiscoverySnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mDiscoverySnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mDiscoverySnapshotMissed.M(1))
			}
		}

		defer func() {
			close(snapshots)
			// we must wait for done before closing the error chan,
			// to avoid sending on close channel.
			done.Wait()
			close(errs)
		}()
		for {
			record := func() { stats.Record(ctx, mDiscoverySnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case upstreamNamespacedList, ok := <-upstreamChan:
				if !ok {
					return
				}
				record()

				namespace := upstreamNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"upstream",
					mDiscoveryResourcesIn,
				)

				// merge lists by namespace
				upstreamsByNamespace[namespace] = upstreamNamespacedList.list
				var upstreamList UpstreamList
				for _, upstreams := range upstreamsByNamespace {
					upstreamList = append(upstreamList, upstreams...)
				}
				currentSnapshot.Upstreams = upstreamList.Sort()
			case kubeNamespaceList, ok := <-kubeNamespaceChan:
				if !ok {
					return
				}
				record()

				skstats.IncrementResourceCount(
					ctx,
					"<all>",
					"kube_namespace",
					mDiscoveryResourcesIn,
				)

				currentSnapshot.Kubenamespaces = kubeNamespaceList
			case secretNamespacedList, ok := <-secretChan:
				if !ok {
					return
				}
				record()

				namespace := secretNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"secret",
					mDiscoveryResourcesIn,
				)

				// merge lists by namespace
				secretsByNamespace[namespace] = secretNamespacedList.list
				var secretList SecretList
				for _, secrets := range secretsByNamespace {
					secretList = append(secretList, secrets...)
				}
				currentSnapshot.Secrets = secretList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}