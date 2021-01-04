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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha3 "kubevirt.io/client-go/api/v1alpha3"
)

// FakeVirtualMachineInstances implements VirtualMachineInstanceInterface
type FakeVirtualMachineInstances struct {
	Fake *FakeKubevirtV1alpha3
	ns   string
}

var virtualmachineinstancesResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1alpha3", Resource: "virtualmachineinstances"}

var virtualmachineinstancesKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1alpha3", Kind: "VirtualMachineInstance"}

// Get takes name of the virtualMachineInstance, and returns the corresponding virtualMachineInstance object, and an error if there is any.
func (c *FakeVirtualMachineInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancesResource, c.ns, name), &v1alpha3.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.VirtualMachineInstance), err
}

// List takes label and field selectors, and returns the list of VirtualMachineInstances that match those selectors.
func (c *FakeVirtualMachineInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.VirtualMachineInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancesResource, virtualmachineinstancesKind, c.ns, opts), &v1alpha3.VirtualMachineInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.VirtualMachineInstanceList{ListMeta: obj.(*v1alpha3.VirtualMachineInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha3.VirtualMachineInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstances.
func (c *FakeVirtualMachineInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineInstance and creates it.  Returns the server's representation of the virtualMachineInstance, and an error, if there is any.
func (c *FakeVirtualMachineInstances) Create(ctx context.Context, virtualMachineInstance *v1alpha3.VirtualMachineInstance, opts v1.CreateOptions) (result *v1alpha3.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &v1alpha3.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.VirtualMachineInstance), err
}

// Update takes the representation of a virtualMachineInstance and updates it. Returns the server's representation of the virtualMachineInstance, and an error, if there is any.
func (c *FakeVirtualMachineInstances) Update(ctx context.Context, virtualMachineInstance *v1alpha3.VirtualMachineInstance, opts v1.UpdateOptions) (result *v1alpha3.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &v1alpha3.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.VirtualMachineInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineInstances) UpdateStatus(ctx context.Context, virtualMachineInstance *v1alpha3.VirtualMachineInstance, opts v1.UpdateOptions) (*v1alpha3.VirtualMachineInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancesResource, "status", c.ns, virtualMachineInstance), &v1alpha3.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.VirtualMachineInstance), err
}

// Delete takes name of the virtualMachineInstance and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancesResource, c.ns, name), &v1alpha3.VirtualMachineInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.VirtualMachineInstanceList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineInstance.
func (c *FakeVirtualMachineInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.VirtualMachineInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha3.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.VirtualMachineInstance), err
}
