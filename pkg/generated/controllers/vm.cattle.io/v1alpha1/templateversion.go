/*
Copyright 2020 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/rancher/harvester/pkg/apis/vm.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type TemplateVersionHandler func(string, *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error)

type TemplateVersionController interface {
	generic.ControllerMeta
	TemplateVersionClient

	OnChange(ctx context.Context, name string, sync TemplateVersionHandler)
	OnRemove(ctx context.Context, name string, sync TemplateVersionHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() TemplateVersionCache
}

type TemplateVersionClient interface {
	Create(*v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error)
	Update(*v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error)
	UpdateStatus(*v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.TemplateVersion, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha1.TemplateVersionList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TemplateVersion, err error)
}

type TemplateVersionCache interface {
	Get(namespace, name string) (*v1alpha1.TemplateVersion, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.TemplateVersion, error)

	AddIndexer(indexName string, indexer TemplateVersionIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha1.TemplateVersion, error)
}

type TemplateVersionIndexer func(obj *v1alpha1.TemplateVersion) ([]string, error)

type templateVersionController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewTemplateVersionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) TemplateVersionController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &templateVersionController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromTemplateVersionHandlerToHandler(sync TemplateVersionHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha1.TemplateVersion
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha1.TemplateVersion))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *templateVersionController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha1.TemplateVersion))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateTemplateVersionDeepCopyOnChange(client TemplateVersionClient, obj *v1alpha1.TemplateVersion, handler func(obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error)) (*v1alpha1.TemplateVersion, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *templateVersionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *templateVersionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *templateVersionController) OnChange(ctx context.Context, name string, sync TemplateVersionHandler) {
	c.AddGenericHandler(ctx, name, FromTemplateVersionHandlerToHandler(sync))
}

func (c *templateVersionController) OnRemove(ctx context.Context, name string, sync TemplateVersionHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromTemplateVersionHandlerToHandler(sync)))
}

func (c *templateVersionController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *templateVersionController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *templateVersionController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *templateVersionController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *templateVersionController) Cache() TemplateVersionCache {
	return &templateVersionCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *templateVersionController) Create(obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error) {
	result := &v1alpha1.TemplateVersion{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *templateVersionController) Update(obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error) {
	result := &v1alpha1.TemplateVersion{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *templateVersionController) UpdateStatus(obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error) {
	result := &v1alpha1.TemplateVersion{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *templateVersionController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *templateVersionController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.TemplateVersion, error) {
	result := &v1alpha1.TemplateVersion{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *templateVersionController) List(namespace string, opts metav1.ListOptions) (*v1alpha1.TemplateVersionList, error) {
	result := &v1alpha1.TemplateVersionList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *templateVersionController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *templateVersionController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha1.TemplateVersion, error) {
	result := &v1alpha1.TemplateVersion{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type templateVersionCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *templateVersionCache) Get(namespace, name string) (*v1alpha1.TemplateVersion, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha1.TemplateVersion), nil
}

func (c *templateVersionCache) List(namespace string, selector labels.Selector) (ret []*v1alpha1.TemplateVersion, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TemplateVersion))
	})

	return ret, err
}

func (c *templateVersionCache) AddIndexer(indexName string, indexer TemplateVersionIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha1.TemplateVersion))
		},
	}))
}

func (c *templateVersionCache) GetByIndex(indexName, key string) (result []*v1alpha1.TemplateVersion, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha1.TemplateVersion, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha1.TemplateVersion))
	}
	return result, nil
}

type TemplateVersionStatusHandler func(obj *v1alpha1.TemplateVersion, status v1alpha1.TemplateVersionStatus) (v1alpha1.TemplateVersionStatus, error)

type TemplateVersionGeneratingHandler func(obj *v1alpha1.TemplateVersion, status v1alpha1.TemplateVersionStatus) ([]runtime.Object, v1alpha1.TemplateVersionStatus, error)

func RegisterTemplateVersionStatusHandler(ctx context.Context, controller TemplateVersionController, condition condition.Cond, name string, handler TemplateVersionStatusHandler) {
	statusHandler := &templateVersionStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromTemplateVersionHandlerToHandler(statusHandler.sync))
}

func RegisterTemplateVersionGeneratingHandler(ctx context.Context, controller TemplateVersionController, apply apply.Apply,
	condition condition.Cond, name string, handler TemplateVersionGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &templateVersionGeneratingHandler{
		TemplateVersionGeneratingHandler: handler,
		apply:                            apply,
		name:                             name,
		gvk:                              controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterTemplateVersionStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type templateVersionStatusHandler struct {
	client    TemplateVersionClient
	condition condition.Cond
	handler   TemplateVersionStatusHandler
}

func (a *templateVersionStatusHandler) sync(key string, obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		var newErr error
		obj.Status = newStatus
		obj, newErr = a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
	}
	return obj, err
}

type templateVersionGeneratingHandler struct {
	TemplateVersionGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *templateVersionGeneratingHandler) Remove(key string, obj *v1alpha1.TemplateVersion) (*v1alpha1.TemplateVersion, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha1.TemplateVersion{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *templateVersionGeneratingHandler) Handle(obj *v1alpha1.TemplateVersion, status v1alpha1.TemplateVersionStatus) (v1alpha1.TemplateVersionStatus, error) {
	objs, newStatus, err := a.TemplateVersionGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
