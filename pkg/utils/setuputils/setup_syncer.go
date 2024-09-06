package setuputils

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector"
	"go.opencensus.io/tag"

	"github.com/solo-io/gloo/pkg/utils"
	"github.com/solo-io/gloo/pkg/utils/settingsutil"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"go.uber.org/zap"
)

var (
	mSetupsRun = utils.MakeSumCounter("gloo.solo.io/setups_run", "The number of times the main setup loop has run")

	mNamespacesWatched      = utils.MakeLastValueCounter("gloo.solo.io/namespaces_watched", "The number of namespaces watched by the gloo controller")
	namespacesWatchedKey, _ = tag.NewKey("namespaces_watched")
)

// tell us how to setup
type SetupFunc func(ctx context.Context,
	kubeCache kube.SharedCache,
	inMemoryCache memory.InMemoryResourceCache,
	settings *v1.Settings,
	identity leaderelector.Identity) error

type SetupSyncer struct {
	settingsRef         *core.ResourceRef
	setupFunc           SetupFunc
	inMemoryCache       memory.InMemoryResourceCache
	identity            leaderelector.Identity
	currentSettingsHash uint64
}

func NewSetupSyncer(settingsRef *core.ResourceRef, setupFunc SetupFunc, identity leaderelector.Identity) *SetupSyncer {
	return &SetupSyncer{
		settingsRef:   settingsRef,
		setupFunc:     setupFunc,
		inMemoryCache: memory.NewInMemoryResourceCache(),
		identity:      identity,
	}
}

func (s *SetupSyncer) Sync(ctx context.Context, snap *v1.SetupSnapshot) error {
	fmt.Println("----------------------- Sync")

	settings, err := snap.Settings.Find(s.settingsRef.Strings())
	if err != nil {
		return errors.Wrapf(err, "finding bootstrap configuration")
	}
	ctx = settingsutil.WithSettings(ctx, settings)

	contextutils.LoggerFrom(ctx).Debugw("received settings snapshot", zap.Any("settings", settings))

	_, err = settingsutil.UpdateNamespacesToWatch(settings, snap.Kubenamespaces)
	if err != nil {
		return err
	}

	watchedNamespaces := settingsutil.GetNamespacesToWatch(settings)
	contextutils.LoggerFrom(ctx).Debugw("received updated list of namespaces to watch", zap.Any("namespaces", watchedNamespaces))

	watchedNamespacesStr := strings.Join(watchedNamespaces, ",")
	utils.Measure(
		ctx,
		mNamespacesWatched,
		int64(len(watchedNamespaces)),
		tag.Insert(namespacesWatchedKey, watchedNamespacesStr),
	)

	var returnErr error
	defer func() {
		// Run this after the function returns to ensure the hash is changed after we successfully sync
		if returnErr == nil {
			s.currentSettingsHash = settings.MustHash()
		}
	}()

	utils.MeasureOne(
		ctx,
		mSetupsRun,
	)

	returnErr = s.setupFunc(ctx, kube.NewKubeCache(ctx), s.inMemoryCache, settings, s.identity)
	return returnErr
}

// ShouldSync compares two snapshots and determines whether a sync is needed based on the following conditions
// 1. The settings CR has changed
// 2. A namespace is added / deleted / modified that changes the namespaces to watch
func (s *SetupSyncer) ShouldSync(ctx context.Context, old, new *v1.SetupSnapshot) bool {
	fmt.Println(("------------------------------ ShouldSync"))
	// Basic sanity checks. Return a true if there is an error to ensure a sync to get into a good state
	if old == nil {
		fmt.Println(("------------------------------ ShouldSync old nil true"))
		return true
	}

	if new == nil {
		fmt.Println(("------------------------------ ShouldSync new nil true"))
		return true
	}

	newSettings, err := new.Settings.Find(s.settingsRef.Strings())
	if err != nil {
		fmt.Println(("------------------------------ ShouldSync new.Settings.Find true"))
		return true
	}

	oldSettings, err := old.Settings.Find(s.settingsRef.Strings())
	if err != nil {
		fmt.Println(("------------------------------ ShouldSync Settings.Find true"))
		return true
	}

	// A sync is triggered if either :
	// 1. The settings CR is changed
	// 2. A namespace is added / deleted / modified

	// 1. Check whether the settings CR is changed
	if s.currentSettingsHash != newSettings.MustHash() {
		fmt.Println(("------------------------------ ShouldSync newSettings.MustHash true"))
		return true
	}

	// 2. Check if a namespace is added / deleted / modified
	// If a namespace was modified, check if it changes the namespaces to watch
	newNamespacesToWatch, err := settingsutil.GenerateNamespacesToWatch(newSettings, new.Kubenamespaces)
	if err != nil {
		fmt.Println(("------------------------------ ShouldSync settingsutil.GenerateNamespacesToWatch true"))
		return true
	}
	namespacesChanged := !slices.Equal(newNamespacesToWatch, settingsutil.GetNamespacesToWatch(oldSettings))
	fmt.Println("------------------------------ ShouldSync namespacesChanged : ", namespacesChanged, settingsutil.GetNamespacesToWatch(oldSettings), newNamespacesToWatch)
	return namespacesChanged
}
