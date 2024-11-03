/*
Copyright The Flagger Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/nholuongut/flagger/pkg/apis/appmesh/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualNodes implements VirtualNodeInterface
type FakeVirtualNodes struct {
	Fake *FakeAppmeshV1beta1
	ns   string
}

var virtualnodesResource = schema.GroupVersionResource{Group: "appmesh.k8s.aws", Version: "v1beta1", Resource: "virtualnodes"}

var virtualnodesKind = schema.GroupVersionKind{Group: "appmesh.k8s.aws", Version: "v1beta1", Kind: "VirtualNode"}

// Get takes name of the virtualNode, and returns the corresponding virtualNode object, and an error if there is any.
func (c *FakeVirtualNodes) Get(name string, options v1.GetOptions) (result *v1beta1.VirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualnodesResource, c.ns, name), &v1beta1.VirtualNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualNode), err
}

// List takes label and field selectors, and returns the list of VirtualNodes that match those selectors.
func (c *FakeVirtualNodes) List(opts v1.ListOptions) (result *v1beta1.VirtualNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualnodesResource, virtualnodesKind, c.ns, opts), &v1beta1.VirtualNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualNodeList{ListMeta: obj.(*v1beta1.VirtualNodeList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualNodes.
func (c *FakeVirtualNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualnodesResource, c.ns, opts))

}

// Create takes the representation of a virtualNode and creates it.  Returns the server's representation of the virtualNode, and an error, if there is any.
func (c *FakeVirtualNodes) Create(virtualNode *v1beta1.VirtualNode) (result *v1beta1.VirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualnodesResource, c.ns, virtualNode), &v1beta1.VirtualNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualNode), err
}

// Update takes the representation of a virtualNode and updates it. Returns the server's representation of the virtualNode, and an error, if there is any.
func (c *FakeVirtualNodes) Update(virtualNode *v1beta1.VirtualNode) (result *v1beta1.VirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualnodesResource, c.ns, virtualNode), &v1beta1.VirtualNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualNodes) UpdateStatus(virtualNode *v1beta1.VirtualNode) (*v1beta1.VirtualNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualnodesResource, "status", c.ns, virtualNode), &v1beta1.VirtualNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualNode), err
}

// Delete takes name of the virtualNode and deletes it. Returns an error if one occurs.
func (c *FakeVirtualNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualnodesResource, c.ns, name), &v1beta1.VirtualNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualnodesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualNodeList{})
	return err
}

// Patch applies the patch and returns the patched virtualNode.
func (c *FakeVirtualNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VirtualNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualnodesResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualNode), err
}
