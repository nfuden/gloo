// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconciler.go -destination mocks/reconciler.go

// The Input Reconciler calls a simple func() error whenever a
// storage event is received for any of:
// * Services
// * Pods
// * Deployments
// * DaemonSets
// * Gateways
// * VirtualServices
// * RouteTables
// * Upstreams
// * UpstreamGroups
// * Settings
// * Proxies
// * AuthConfigs
// * RateLimitConfigs
// for a given cluster or set of clusters.
//
// Input Reconcilers can be be constructed from either a single Manager (watch events in a single cluster)
// or a ClusterWatcher (watch events in multiple clusters).
package input

import (
	"context"
	"time"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	v1 "k8s.io/api/core/v1"

	apps_v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	apps_v1 "k8s.io/api/apps/v1"

	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	gateway_solo_io_v1_controllers "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1/controller"

	gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	gloo_solo_io_v1_controllers "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/controller"

	enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	enterprise_gloo_solo_io_v1_controllers "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1/controller"

	ratelimit_api_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	ratelimit_api_solo_io_v1alpha1_controllers "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1/controller"
)

// the multiClusterReconciler reconciles events for input resources across clusters
// this private interface is used to ensure that the generated struct implements the intended functions
type multiClusterReconciler interface {
	v1_controllers.MulticlusterServiceReconciler
	v1_controllers.MulticlusterPodReconciler

	apps_v1_controllers.MulticlusterDeploymentReconciler
	apps_v1_controllers.MulticlusterDaemonSetReconciler

	gateway_solo_io_v1_controllers.MulticlusterGatewayReconciler
	gateway_solo_io_v1_controllers.MulticlusterVirtualServiceReconciler
	gateway_solo_io_v1_controllers.MulticlusterRouteTableReconciler

	gloo_solo_io_v1_controllers.MulticlusterUpstreamReconciler
	gloo_solo_io_v1_controllers.MulticlusterUpstreamGroupReconciler
	gloo_solo_io_v1_controllers.MulticlusterSettingsReconciler
	gloo_solo_io_v1_controllers.MulticlusterProxyReconciler

	enterprise_gloo_solo_io_v1_controllers.MulticlusterAuthConfigReconciler

	ratelimit_api_solo_io_v1alpha1_controllers.MulticlusterRateLimitConfigReconciler
}

var _ multiClusterReconciler = &multiClusterReconcilerImpl{}

type multiClusterReconcilerImpl struct {
	base input.InputReconciler
}

// Options for reconciling a snapshot
type ReconcileOptions struct {

	// Options for reconciling Services
	Services reconcile.Options
	// Options for reconciling Pods
	Pods reconcile.Options

	// Options for reconciling Deployments
	Deployments reconcile.Options
	// Options for reconciling DaemonSets
	DaemonSets reconcile.Options

	// Options for reconciling Gateways
	Gateways reconcile.Options
	// Options for reconciling VirtualServices
	VirtualServices reconcile.Options
	// Options for reconciling RouteTables
	RouteTables reconcile.Options

	// Options for reconciling Upstreams
	Upstreams reconcile.Options
	// Options for reconciling UpstreamGroups
	UpstreamGroups reconcile.Options
	// Options for reconciling Settings
	Settings reconcile.Options
	// Options for reconciling Proxies
	Proxies reconcile.Options

	// Options for reconciling AuthConfigs
	AuthConfigs reconcile.Options

	// Options for reconciling RateLimitConfigs
	RateLimitConfigs reconcile.Options
}

// register the reconcile func with the cluster watcher
// the reconcileInterval, if greater than 0, will limit the number of reconciles
// to one per interval.
func RegisterMultiClusterReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	reconcileFunc input.MultiClusterReconcileFunc,
	reconcileInterval time.Duration,
	options ReconcileOptions,
	predicates ...predicate.Predicate,
) input.InputReconciler {

	base := input.NewInputReconciler(
		ctx,
		reconcileFunc,
		nil,
		reconcileInterval,
	)

	r := &multiClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	v1_controllers.NewMulticlusterServiceReconcileLoop("Service", clusters, options.Services).AddMulticlusterServiceReconciler(ctx, r, predicates...)

	v1_controllers.NewMulticlusterPodReconcileLoop("Pod", clusters, options.Pods).AddMulticlusterPodReconciler(ctx, r, predicates...)

	apps_v1_controllers.NewMulticlusterDeploymentReconcileLoop("Deployment", clusters, options.Deployments).AddMulticlusterDeploymentReconciler(ctx, r, predicates...)

	apps_v1_controllers.NewMulticlusterDaemonSetReconcileLoop("DaemonSet", clusters, options.DaemonSets).AddMulticlusterDaemonSetReconciler(ctx, r, predicates...)

	gateway_solo_io_v1_controllers.NewMulticlusterGatewayReconcileLoop("Gateway", clusters, options.Gateways).AddMulticlusterGatewayReconciler(ctx, r, predicates...)

	gateway_solo_io_v1_controllers.NewMulticlusterVirtualServiceReconcileLoop("VirtualService", clusters, options.VirtualServices).AddMulticlusterVirtualServiceReconciler(ctx, r, predicates...)

	gateway_solo_io_v1_controllers.NewMulticlusterRouteTableReconcileLoop("RouteTable", clusters, options.RouteTables).AddMulticlusterRouteTableReconciler(ctx, r, predicates...)

	gloo_solo_io_v1_controllers.NewMulticlusterUpstreamReconcileLoop("Upstream", clusters, options.Upstreams).AddMulticlusterUpstreamReconciler(ctx, r, predicates...)

	gloo_solo_io_v1_controllers.NewMulticlusterUpstreamGroupReconcileLoop("UpstreamGroup", clusters, options.UpstreamGroups).AddMulticlusterUpstreamGroupReconciler(ctx, r, predicates...)

	gloo_solo_io_v1_controllers.NewMulticlusterSettingsReconcileLoop("Settings", clusters, options.Settings).AddMulticlusterSettingsReconciler(ctx, r, predicates...)

	gloo_solo_io_v1_controllers.NewMulticlusterProxyReconcileLoop("Proxy", clusters, options.Proxies).AddMulticlusterProxyReconciler(ctx, r, predicates...)

	enterprise_gloo_solo_io_v1_controllers.NewMulticlusterAuthConfigReconcileLoop("AuthConfig", clusters, options.AuthConfigs).AddMulticlusterAuthConfigReconciler(ctx, r, predicates...)

	ratelimit_api_solo_io_v1alpha1_controllers.NewMulticlusterRateLimitConfigReconcileLoop("RateLimitConfig", clusters, options.RateLimitConfigs).AddMulticlusterRateLimitConfigReconciler(ctx, r, predicates...)
	return r.base
}

func (r *multiClusterReconcilerImpl) ReconcileService(clusterName string, obj *v1.Service) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcilePod(clusterName string, obj *v1.Pod) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcilePodDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileDeployment(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileDeploymentDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileDaemonSet(clusterName string, obj *apps_v1.DaemonSet) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileDaemonSetDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileGateway(clusterName string, obj *gateway_solo_io_v1.Gateway) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileGatewayDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualService(clusterName string, obj *gateway_solo_io_v1.VirtualService) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileRouteTable(clusterName string, obj *gateway_solo_io_v1.RouteTable) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileRouteTableDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileUpstream(clusterName string, obj *gloo_solo_io_v1.Upstream) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileUpstreamDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileUpstreamGroup(clusterName string, obj *gloo_solo_io_v1.UpstreamGroup) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileUpstreamGroupDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileSettings(clusterName string, obj *gloo_solo_io_v1.Settings) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileSettingsDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileProxy(clusterName string, obj *gloo_solo_io_v1.Proxy) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileProxyDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileAuthConfig(clusterName string, obj *enterprise_gloo_solo_io_v1.AuthConfig) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileAuthConfigDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileRateLimitConfig(clusterName string, obj *ratelimit_api_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileRateLimitConfigDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

// the singleClusterReconciler reconciles events for input resources across clusters
// this private interface is used to ensure that the generated struct implements the intended functions
type singleClusterReconciler interface {
	v1_controllers.ServiceReconciler
	v1_controllers.PodReconciler

	apps_v1_controllers.DeploymentReconciler
	apps_v1_controllers.DaemonSetReconciler

	gateway_solo_io_v1_controllers.GatewayReconciler
	gateway_solo_io_v1_controllers.VirtualServiceReconciler
	gateway_solo_io_v1_controllers.RouteTableReconciler

	gloo_solo_io_v1_controllers.UpstreamReconciler
	gloo_solo_io_v1_controllers.UpstreamGroupReconciler
	gloo_solo_io_v1_controllers.SettingsReconciler
	gloo_solo_io_v1_controllers.ProxyReconciler

	enterprise_gloo_solo_io_v1_controllers.AuthConfigReconciler

	ratelimit_api_solo_io_v1alpha1_controllers.RateLimitConfigReconciler
}

var _ singleClusterReconciler = &singleClusterReconcilerImpl{}

type singleClusterReconcilerImpl struct {
	base input.InputReconciler
}

// register the reconcile func with the manager
// the reconcileInterval, if greater than 0, will limit the number of reconciles
// to one per interval.
func RegisterSingleClusterReconciler(
	ctx context.Context,
	mgr manager.Manager,
	reconcileFunc input.SingleClusterReconcileFunc,
	reconcileInterval time.Duration,
	options reconcile.Options,
	predicates ...predicate.Predicate,
) (input.InputReconciler, error) {

	base := input.NewInputReconciler(
		ctx,
		nil,
		reconcileFunc,
		reconcileInterval,
	)

	r := &singleClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	if err := v1_controllers.NewServiceReconcileLoop("Service", mgr, options).RunServiceReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := v1_controllers.NewPodReconcileLoop("Pod", mgr, options).RunPodReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	if err := apps_v1_controllers.NewDeploymentReconcileLoop("Deployment", mgr, options).RunDeploymentReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := apps_v1_controllers.NewDaemonSetReconcileLoop("DaemonSet", mgr, options).RunDaemonSetReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	if err := gateway_solo_io_v1_controllers.NewGatewayReconcileLoop("Gateway", mgr, options).RunGatewayReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := gateway_solo_io_v1_controllers.NewVirtualServiceReconcileLoop("VirtualService", mgr, options).RunVirtualServiceReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := gateway_solo_io_v1_controllers.NewRouteTableReconcileLoop("RouteTable", mgr, options).RunRouteTableReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	if err := gloo_solo_io_v1_controllers.NewUpstreamReconcileLoop("Upstream", mgr, options).RunUpstreamReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := gloo_solo_io_v1_controllers.NewUpstreamGroupReconcileLoop("UpstreamGroup", mgr, options).RunUpstreamGroupReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := gloo_solo_io_v1_controllers.NewSettingsReconcileLoop("Settings", mgr, options).RunSettingsReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}
	if err := gloo_solo_io_v1_controllers.NewProxyReconcileLoop("Proxy", mgr, options).RunProxyReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	if err := enterprise_gloo_solo_io_v1_controllers.NewAuthConfigReconcileLoop("AuthConfig", mgr, options).RunAuthConfigReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	if err := ratelimit_api_solo_io_v1alpha1_controllers.NewRateLimitConfigReconcileLoop("RateLimitConfig", mgr, options).RunRateLimitConfigReconciler(ctx, r, predicates...); err != nil {
		return nil, err
	}

	return r.base, nil
}

func (r *singleClusterReconcilerImpl) ReconcileService(obj *v1.Service) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileServiceDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcilePod(obj *v1.Pod) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcilePodDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileDeployment(obj *apps_v1.Deployment) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileDeploymentDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileDaemonSet(obj *apps_v1.DaemonSet) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileDaemonSetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileGateway(obj *gateway_solo_io_v1.Gateway) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileGatewayDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualService(obj *gateway_solo_io_v1.VirtualService) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualServiceDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileRouteTable(obj *gateway_solo_io_v1.RouteTable) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileRouteTableDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileUpstream(obj *gloo_solo_io_v1.Upstream) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileUpstreamDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileUpstreamGroup(obj *gloo_solo_io_v1.UpstreamGroup) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileUpstreamGroupDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileSettings(obj *gloo_solo_io_v1.Settings) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileSettingsDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileProxy(obj *gloo_solo_io_v1.Proxy) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileProxyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileAuthConfig(obj *enterprise_gloo_solo_io_v1.AuthConfig) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileAuthConfigDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileRateLimitConfig(obj *ratelimit_api_solo_io_v1alpha1.RateLimitConfig) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileRateLimitConfigDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}
