// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the multicluster.solo.io/v1alpha1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the multicluster.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
	MultiClusterRoles() MultiClusterRoleClient
	// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
	MultiClusterRoleBindings() MultiClusterRoleBindingClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) MultiClusterRoles() MultiClusterRoleClient {
	return NewMultiClusterRoleClient(c.client)
}

// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) MultiClusterRoleBindings() MultiClusterRoleBindingClient {
	return NewMultiClusterRoleBindingClient(c.client)
}

// Reader knows how to read and list MultiClusterRoles.
type MultiClusterRoleReader interface {
	// Get retrieves a MultiClusterRole for the given object key
	GetMultiClusterRole(ctx context.Context, key client.ObjectKey) (*MultiClusterRole, error)

	// List retrieves list of MultiClusterRoles for a given namespace and list options.
	ListMultiClusterRole(ctx context.Context, opts ...client.ListOption) (*MultiClusterRoleList, error)
}

// MultiClusterRoleTransitionFunction instructs the MultiClusterRoleWriter how to transition between an existing
// MultiClusterRole object and a desired on an Upsert
type MultiClusterRoleTransitionFunction func(existing, desired *MultiClusterRole) error

// Writer knows how to create, delete, and update MultiClusterRoles.
type MultiClusterRoleWriter interface {
	// Create saves the MultiClusterRole object.
	CreateMultiClusterRole(ctx context.Context, obj *MultiClusterRole, opts ...client.CreateOption) error

	// Delete deletes the MultiClusterRole object.
	DeleteMultiClusterRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given MultiClusterRole object.
	UpdateMultiClusterRole(ctx context.Context, obj *MultiClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given MultiClusterRole object.
	PatchMultiClusterRole(ctx context.Context, obj *MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all MultiClusterRole objects matching the given options.
	DeleteAllOfMultiClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the MultiClusterRole object.
	UpsertMultiClusterRole(ctx context.Context, obj *MultiClusterRole, transitionFuncs ...MultiClusterRoleTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a MultiClusterRole object.
type MultiClusterRoleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given MultiClusterRole object.
	UpdateMultiClusterRoleStatus(ctx context.Context, obj *MultiClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given MultiClusterRole object's subresource.
	PatchMultiClusterRoleStatus(ctx context.Context, obj *MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on MultiClusterRoles.
type MultiClusterRoleClient interface {
	MultiClusterRoleReader
	MultiClusterRoleWriter
	MultiClusterRoleStatusWriter
}

type multiClusterRoleClient struct {
	client client.Client
}

func NewMultiClusterRoleClient(client client.Client) *multiClusterRoleClient {
	return &multiClusterRoleClient{client: client}
}

func (c *multiClusterRoleClient) GetMultiClusterRole(ctx context.Context, key client.ObjectKey) (*MultiClusterRole, error) {
	obj := &MultiClusterRole{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *multiClusterRoleClient) ListMultiClusterRole(ctx context.Context, opts ...client.ListOption) (*MultiClusterRoleList, error) {
	list := &MultiClusterRoleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *multiClusterRoleClient) CreateMultiClusterRole(ctx context.Context, obj *MultiClusterRole, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *multiClusterRoleClient) DeleteMultiClusterRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &MultiClusterRole{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *multiClusterRoleClient) UpdateMultiClusterRole(ctx context.Context, obj *MultiClusterRole, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *multiClusterRoleClient) PatchMultiClusterRole(ctx context.Context, obj *MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *multiClusterRoleClient) DeleteAllOfMultiClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &MultiClusterRole{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *multiClusterRoleClient) UpsertMultiClusterRole(ctx context.Context, obj *MultiClusterRole, transitionFuncs ...MultiClusterRoleTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*MultiClusterRole), desired.(*MultiClusterRole)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *multiClusterRoleClient) UpdateMultiClusterRoleStatus(ctx context.Context, obj *MultiClusterRole, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *multiClusterRoleClient) PatchMultiClusterRoleStatus(ctx context.Context, obj *MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides MultiClusterRoleClients for multiple clusters.
type MulticlusterMultiClusterRoleClient interface {
	// Cluster returns a MultiClusterRoleClient for the given cluster
	Cluster(cluster string) (MultiClusterRoleClient, error)
}

type multiclusterMultiClusterRoleClient struct {
	client multicluster.Client
}

func NewMulticlusterMultiClusterRoleClient(client multicluster.Client) MulticlusterMultiClusterRoleClient {
	return &multiclusterMultiClusterRoleClient{client: client}
}

func (m *multiclusterMultiClusterRoleClient) Cluster(cluster string) (MultiClusterRoleClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewMultiClusterRoleClient(client), nil
}

// Reader knows how to read and list MultiClusterRoleBindings.
type MultiClusterRoleBindingReader interface {
	// Get retrieves a MultiClusterRoleBinding for the given object key
	GetMultiClusterRoleBinding(ctx context.Context, key client.ObjectKey) (*MultiClusterRoleBinding, error)

	// List retrieves list of MultiClusterRoleBindings for a given namespace and list options.
	ListMultiClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*MultiClusterRoleBindingList, error)
}

// MultiClusterRoleBindingTransitionFunction instructs the MultiClusterRoleBindingWriter how to transition between an existing
// MultiClusterRoleBinding object and a desired on an Upsert
type MultiClusterRoleBindingTransitionFunction func(existing, desired *MultiClusterRoleBinding) error

// Writer knows how to create, delete, and update MultiClusterRoleBindings.
type MultiClusterRoleBindingWriter interface {
	// Create saves the MultiClusterRoleBinding object.
	CreateMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.CreateOption) error

	// Delete deletes the MultiClusterRoleBinding object.
	DeleteMultiClusterRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given MultiClusterRoleBinding object.
	UpdateMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given MultiClusterRoleBinding object.
	PatchMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all MultiClusterRoleBinding objects matching the given options.
	DeleteAllOfMultiClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the MultiClusterRoleBinding object.
	UpsertMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, transitionFuncs ...MultiClusterRoleBindingTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a MultiClusterRoleBinding object.
type MultiClusterRoleBindingStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given MultiClusterRoleBinding object.
	UpdateMultiClusterRoleBindingStatus(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given MultiClusterRoleBinding object's subresource.
	PatchMultiClusterRoleBindingStatus(ctx context.Context, obj *MultiClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on MultiClusterRoleBindings.
type MultiClusterRoleBindingClient interface {
	MultiClusterRoleBindingReader
	MultiClusterRoleBindingWriter
	MultiClusterRoleBindingStatusWriter
}

type multiClusterRoleBindingClient struct {
	client client.Client
}

func NewMultiClusterRoleBindingClient(client client.Client) *multiClusterRoleBindingClient {
	return &multiClusterRoleBindingClient{client: client}
}

func (c *multiClusterRoleBindingClient) GetMultiClusterRoleBinding(ctx context.Context, key client.ObjectKey) (*MultiClusterRoleBinding, error) {
	obj := &MultiClusterRoleBinding{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *multiClusterRoleBindingClient) ListMultiClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*MultiClusterRoleBindingList, error) {
	list := &MultiClusterRoleBindingList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *multiClusterRoleBindingClient) CreateMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *multiClusterRoleBindingClient) DeleteMultiClusterRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &MultiClusterRoleBinding{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *multiClusterRoleBindingClient) UpdateMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *multiClusterRoleBindingClient) PatchMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *multiClusterRoleBindingClient) DeleteAllOfMultiClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &MultiClusterRoleBinding{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *multiClusterRoleBindingClient) UpsertMultiClusterRoleBinding(ctx context.Context, obj *MultiClusterRoleBinding, transitionFuncs ...MultiClusterRoleBindingTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*MultiClusterRoleBinding), desired.(*MultiClusterRoleBinding)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *multiClusterRoleBindingClient) UpdateMultiClusterRoleBindingStatus(ctx context.Context, obj *MultiClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *multiClusterRoleBindingClient) PatchMultiClusterRoleBindingStatus(ctx context.Context, obj *MultiClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides MultiClusterRoleBindingClients for multiple clusters.
type MulticlusterMultiClusterRoleBindingClient interface {
	// Cluster returns a MultiClusterRoleBindingClient for the given cluster
	Cluster(cluster string) (MultiClusterRoleBindingClient, error)
}

type multiclusterMultiClusterRoleBindingClient struct {
	client multicluster.Client
}

func NewMulticlusterMultiClusterRoleBindingClient(client multicluster.Client) MulticlusterMultiClusterRoleBindingClient {
	return &multiclusterMultiClusterRoleBindingClient{client: client}
}

func (m *multiclusterMultiClusterRoleBindingClient) Cluster(cluster string) (MultiClusterRoleBindingClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewMultiClusterRoleBindingClient(client), nil
}