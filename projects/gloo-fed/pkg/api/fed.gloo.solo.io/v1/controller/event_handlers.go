// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	fed_gloo_solo_io_v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the FederatedUpstream Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedUpstreamEventHandler interface {
	CreateFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
	UpdateFederatedUpstream(old, new *fed_gloo_solo_io_v1.FederatedUpstream) error
	DeleteFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
	GenericFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
}

type FederatedUpstreamEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
	OnUpdate  func(old, new *fed_gloo_solo_io_v1.FederatedUpstream) error
	OnDelete  func(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
	OnGeneric func(obj *fed_gloo_solo_io_v1.FederatedUpstream) error
}

func (f *FederatedUpstreamEventHandlerFuncs) CreateFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedUpstreamEventHandlerFuncs) DeleteFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedUpstreamEventHandlerFuncs) UpdateFederatedUpstream(objOld, objNew *fed_gloo_solo_io_v1.FederatedUpstream) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedUpstreamEventHandlerFuncs) GenericFederatedUpstream(obj *fed_gloo_solo_io_v1.FederatedUpstream) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedUpstreamEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedUpstreamEventHandler, predicates ...predicate.Predicate) error
}

type federatedUpstreamEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedUpstreamEventWatcher(name string, mgr manager.Manager) FederatedUpstreamEventWatcher {
	return &federatedUpstreamEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gloo_solo_io_v1.FederatedUpstream{}),
	}
}

func (c *federatedUpstreamEventWatcher) AddEventHandler(ctx context.Context, h FederatedUpstreamEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedUpstreamHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedUpstreamHandler implements a generic events.EventHandler
type genericFederatedUpstreamHandler struct {
	handler FederatedUpstreamEventHandler
}

func (h genericFederatedUpstreamHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstream handler received event for %T", object)
	}
	return h.handler.CreateFederatedUpstream(obj)
}

func (h genericFederatedUpstreamHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstream handler received event for %T", object)
	}
	return h.handler.DeleteFederatedUpstream(obj)
}

func (h genericFederatedUpstreamHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstream handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstream handler received event for %T", new)
	}
	return h.handler.UpdateFederatedUpstream(objOld, objNew)
}

func (h genericFederatedUpstreamHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstream handler received event for %T", object)
	}
	return h.handler.GenericFederatedUpstream(obj)
}

// Handle events for the FederatedUpstreamGroup Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedUpstreamGroupEventHandler interface {
	CreateFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	UpdateFederatedUpstreamGroup(old, new *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	DeleteFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	GenericFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
}

type FederatedUpstreamGroupEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	OnUpdate  func(old, new *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	OnDelete  func(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
	OnGeneric func(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error
}

func (f *FederatedUpstreamGroupEventHandlerFuncs) CreateFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedUpstreamGroupEventHandlerFuncs) DeleteFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedUpstreamGroupEventHandlerFuncs) UpdateFederatedUpstreamGroup(objOld, objNew *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedUpstreamGroupEventHandlerFuncs) GenericFederatedUpstreamGroup(obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedUpstreamGroupEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedUpstreamGroupEventHandler, predicates ...predicate.Predicate) error
}

type federatedUpstreamGroupEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedUpstreamGroupEventWatcher(name string, mgr manager.Manager) FederatedUpstreamGroupEventWatcher {
	return &federatedUpstreamGroupEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gloo_solo_io_v1.FederatedUpstreamGroup{}),
	}
}

func (c *federatedUpstreamGroupEventWatcher) AddEventHandler(ctx context.Context, h FederatedUpstreamGroupEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedUpstreamGroupHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedUpstreamGroupHandler implements a generic events.EventHandler
type genericFederatedUpstreamGroupHandler struct {
	handler FederatedUpstreamGroupEventHandler
}

func (h genericFederatedUpstreamGroupHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", object)
	}
	return h.handler.CreateFederatedUpstreamGroup(obj)
}

func (h genericFederatedUpstreamGroupHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", object)
	}
	return h.handler.DeleteFederatedUpstreamGroup(obj)
}

func (h genericFederatedUpstreamGroupHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", new)
	}
	return h.handler.UpdateFederatedUpstreamGroup(objOld, objNew)
}

func (h genericFederatedUpstreamGroupHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", object)
	}
	return h.handler.GenericFederatedUpstreamGroup(obj)
}

// Handle events for the FederatedSettings Resource
// DEPRECATED: Prefer reconciler pattern.
type FederatedSettingsEventHandler interface {
	CreateFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error
	UpdateFederatedSettings(old, new *fed_gloo_solo_io_v1.FederatedSettings) error
	DeleteFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error
	GenericFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error
}

type FederatedSettingsEventHandlerFuncs struct {
	OnCreate  func(obj *fed_gloo_solo_io_v1.FederatedSettings) error
	OnUpdate  func(old, new *fed_gloo_solo_io_v1.FederatedSettings) error
	OnDelete  func(obj *fed_gloo_solo_io_v1.FederatedSettings) error
	OnGeneric func(obj *fed_gloo_solo_io_v1.FederatedSettings) error
}

func (f *FederatedSettingsEventHandlerFuncs) CreateFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *FederatedSettingsEventHandlerFuncs) DeleteFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *FederatedSettingsEventHandlerFuncs) UpdateFederatedSettings(objOld, objNew *fed_gloo_solo_io_v1.FederatedSettings) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *FederatedSettingsEventHandlerFuncs) GenericFederatedSettings(obj *fed_gloo_solo_io_v1.FederatedSettings) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type FederatedSettingsEventWatcher interface {
	AddEventHandler(ctx context.Context, h FederatedSettingsEventHandler, predicates ...predicate.Predicate) error
}

type federatedSettingsEventWatcher struct {
	watcher events.EventWatcher
}

func NewFederatedSettingsEventWatcher(name string, mgr manager.Manager) FederatedSettingsEventWatcher {
	return &federatedSettingsEventWatcher{
		watcher: events.NewWatcher(name, mgr, &fed_gloo_solo_io_v1.FederatedSettings{}),
	}
}

func (c *federatedSettingsEventWatcher) AddEventHandler(ctx context.Context, h FederatedSettingsEventHandler, predicates ...predicate.Predicate) error {
	handler := genericFederatedSettingsHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericFederatedSettingsHandler implements a generic events.EventHandler
type genericFederatedSettingsHandler struct {
	handler FederatedSettingsEventHandler
}

func (h genericFederatedSettingsHandler) Create(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return errors.Errorf("internal error: FederatedSettings handler received event for %T", object)
	}
	return h.handler.CreateFederatedSettings(obj)
}

func (h genericFederatedSettingsHandler) Delete(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return errors.Errorf("internal error: FederatedSettings handler received event for %T", object)
	}
	return h.handler.DeleteFederatedSettings(obj)
}

func (h genericFederatedSettingsHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return errors.Errorf("internal error: FederatedSettings handler received event for %T", old)
	}
	objNew, ok := new.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return errors.Errorf("internal error: FederatedSettings handler received event for %T", new)
	}
	return h.handler.UpdateFederatedSettings(objOld, objNew)
}

func (h genericFederatedSettingsHandler) Generic(object client.Object) error {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return errors.Errorf("internal error: FederatedSettings handler received event for %T", object)
	}
	return h.handler.GenericFederatedSettings(obj)
}
