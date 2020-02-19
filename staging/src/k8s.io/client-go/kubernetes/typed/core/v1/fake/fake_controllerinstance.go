/*
Copyright 2020 Authors of Arktos.

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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeControllerInstances implements ControllerInstanceInterface
type FakeControllerInstances struct {
	Fake *FakeCoreV1
}

var controllerinstancesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "controllerinstances"}

var controllerinstancesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ControllerInstance"}

// Get takes name of the controllerInstance, and returns the corresponding controllerInstance object, and an error if there is any.
func (c *FakeControllerInstances) Get(name string, options v1.GetOptions) (result *corev1.ControllerInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(controllerinstancesResource, name), &corev1.ControllerInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.ControllerInstance), err
}

// List takes label and field selectors, and returns the list of ControllerInstances that match those selectors.
func (c *FakeControllerInstances) List(opts v1.ListOptions) (result *corev1.ControllerInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(controllerinstancesResource, controllerinstancesKind, opts), &corev1.ControllerInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ControllerInstanceList{ListMeta: obj.(*corev1.ControllerInstanceList).ListMeta}
	for _, item := range obj.(*corev1.ControllerInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested controllerInstances.
func (c *FakeControllerInstances) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewRootWatchAction(controllerinstancesResource, opts))
	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a controllerInstance and creates it.  Returns the server's representation of the controllerInstance, and an error, if there is any.
func (c *FakeControllerInstances) Create(controllerInstance *corev1.ControllerInstance) (result *corev1.ControllerInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(controllerinstancesResource, controllerInstance), &corev1.ControllerInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.ControllerInstance), err
}

// Update takes the representation of a controllerInstance and updates it. Returns the server's representation of the controllerInstance, and an error, if there is any.
func (c *FakeControllerInstances) Update(controllerInstance *corev1.ControllerInstance) (result *corev1.ControllerInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(controllerinstancesResource, controllerInstance), &corev1.ControllerInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.ControllerInstance), err
}

// Delete takes name of the controllerInstance and deletes it. Returns an error if one occurs.
func (c *FakeControllerInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(controllerinstancesResource, name), &corev1.ControllerInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeControllerInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {

	action := testing.NewRootDeleteCollectionAction(controllerinstancesResource, listOptions)
	_, err := c.Fake.Invokes(action, &corev1.ControllerInstanceList{})
	return err
}

// Patch applies the patch and returns the patched controllerInstance.
func (c *FakeControllerInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.ControllerInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(controllerinstancesResource, name, pt, data, subresources...), &corev1.ControllerInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.ControllerInstance), err
}
