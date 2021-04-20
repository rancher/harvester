/*
Copyright 2021 Rancher Labs, Inc.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/rancher/harvester/pkg/apis/harvesterhci.io/v1beta1"
	scheme "github.com/rancher/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachineTemplatesGetter has a method to return a VirtualMachineTemplateInterface.
// A group's client should implement this interface.
type VirtualMachineTemplatesGetter interface {
	VirtualMachineTemplates(namespace string) VirtualMachineTemplateInterface
}

// VirtualMachineTemplateInterface has methods to work with VirtualMachineTemplate resources.
type VirtualMachineTemplateInterface interface {
	Create(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.CreateOptions) (*v1beta1.VirtualMachineTemplate, error)
	Update(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplate, error)
	UpdateStatus(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplate, err error)
	VirtualMachineTemplateExpansion
}

// virtualMachineTemplates implements VirtualMachineTemplateInterface
type virtualMachineTemplates struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineTemplates returns a VirtualMachineTemplates
func newVirtualMachineTemplates(c *HarvesterhciV1beta1Client, namespace string) *virtualMachineTemplates {
	return &virtualMachineTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineTemplate, and returns the corresponding virtualMachineTemplate object, and an error if there is any.
func (c *virtualMachineTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineTemplates that match those selectors.
func (c *virtualMachineTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachineTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineTemplates.
func (c *virtualMachineTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineTemplate and creates it.  Returns the server's representation of the virtualMachineTemplate, and an error, if there is any.
func (c *virtualMachineTemplates) Create(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.CreateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineTemplate and updates it. Returns the server's representation of the virtualMachineTemplate, and an error, if there is any.
func (c *virtualMachineTemplates) Update(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(virtualMachineTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineTemplates) UpdateStatus(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(virtualMachineTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineTemplate and deletes it. Returns an error if one occurs.
func (c *virtualMachineTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineTemplate.
func (c *virtualMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplate, err error) {
	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}