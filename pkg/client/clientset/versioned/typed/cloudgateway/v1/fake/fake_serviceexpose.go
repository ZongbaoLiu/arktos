/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	cloudgatewayv1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
)

// FakeServiceExposes implements ServiceExposeInterface
type FakeServiceExposes struct {
	Fake *FakeCloudgatewayV1
	ns   string
	te   string
}

var serviceexposesResource = schema.GroupVersionResource{Group: "cloudgateway.arktos.io", Version: "v1", Resource: "serviceexposes"}

var serviceexposesKind = schema.GroupVersionKind{Group: "cloudgateway.arktos.io", Version: "v1", Kind: "ServiceExpose"}

// Get takes name of the serviceExpose, and returns the corresponding serviceExpose object, and an error if there is any.
func (c *FakeServiceExposes) Get(name string, options v1.GetOptions) (result *cloudgatewayv1.ServiceExpose, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(serviceexposesResource, c.ns, name, c.te), &cloudgatewayv1.ServiceExpose{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.ServiceExpose), err
}

// List takes label and field selectors, and returns the list of ServiceExposes that match those selectors.
func (c *FakeServiceExposes) List(opts v1.ListOptions) (result *cloudgatewayv1.ServiceExposeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(serviceexposesResource, serviceexposesKind, c.ns, opts, c.te), &cloudgatewayv1.ServiceExposeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudgatewayv1.ServiceExposeList{ListMeta: obj.(*cloudgatewayv1.ServiceExposeList).ListMeta}
	for _, item := range obj.(*cloudgatewayv1.ServiceExposeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested serviceExposes.
func (c *FakeServiceExposes) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(serviceexposesResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a serviceExpose and creates it.  Returns the server's representation of the serviceExpose, and an error, if there is any.
func (c *FakeServiceExposes) Create(serviceExpose *cloudgatewayv1.ServiceExpose) (result *cloudgatewayv1.ServiceExpose, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(serviceexposesResource, c.ns, serviceExpose, c.te), &cloudgatewayv1.ServiceExpose{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.ServiceExpose), err
}

// Update takes the representation of a serviceExpose and updates it. Returns the server's representation of the serviceExpose, and an error, if there is any.
func (c *FakeServiceExposes) Update(serviceExpose *cloudgatewayv1.ServiceExpose) (result *cloudgatewayv1.ServiceExpose, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(serviceexposesResource, c.ns, serviceExpose, c.te), &cloudgatewayv1.ServiceExpose{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.ServiceExpose), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceExposes) UpdateStatus(serviceExpose *cloudgatewayv1.ServiceExpose) (*cloudgatewayv1.ServiceExpose, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(serviceexposesResource, "status", c.ns, serviceExpose, c.te), &cloudgatewayv1.ServiceExpose{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudgatewayv1.ServiceExpose), err
}

// Delete takes name of the serviceExpose and deletes it. Returns an error if one occurs.
func (c *FakeServiceExposes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(serviceexposesResource, c.ns, name, c.te), &cloudgatewayv1.ServiceExpose{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceExposes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(serviceexposesResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &cloudgatewayv1.ServiceExposeList{})
	return err
}

// Patch applies the patch and returns the patched serviceExpose.
func (c *FakeServiceExposes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudgatewayv1.ServiceExpose, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(serviceexposesResource, c.te, c.ns, name, pt, data, subresources...), &cloudgatewayv1.ServiceExpose{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.ServiceExpose), err
}
