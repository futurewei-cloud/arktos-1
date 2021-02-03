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
	examplev1 "k8s.io/code-generator/_examples/apiserver/apis/example/v1"
)

// FakeTestTypes implements TestTypeInterface
type FakeTestTypes struct {
	Fake *FakeExampleV1
	ns   string
	te   string
}

var testtypesResource = schema.GroupVersionResource{Group: "example.apiserver.code-generator.k8s.io", Version: "v1", Resource: "testtypes"}

var testtypesKind = schema.GroupVersionKind{Group: "example.apiserver.code-generator.k8s.io", Version: "v1", Kind: "TestType"}

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *FakeTestTypes) Get(name string, options v1.GetOptions) (result *examplev1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(testtypesResource, c.ns, name, c.te), &examplev1.TestType{})

	if obj == nil {
		return nil, err
	}

	return obj.(*examplev1.TestType), err
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *FakeTestTypes) List(opts v1.ListOptions) (result *examplev1.TestTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(testtypesResource, testtypesKind, c.ns, opts, c.te), &examplev1.TestTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &examplev1.TestTypeList{ListMeta: obj.(*examplev1.TestTypeList).ListMeta}
	for _, item := range obj.(*examplev1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested testTypes.
func (c *FakeTestTypes) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(testtypesResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Create(testType *examplev1.TestType) (result *examplev1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(testtypesResource, c.ns, testType, c.te), &examplev1.TestType{})

	if obj == nil {
		return nil, err
	}

	return obj.(*examplev1.TestType), err
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Update(testType *examplev1.TestType) (result *examplev1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(testtypesResource, c.ns, testType, c.te), &examplev1.TestType{})

	if obj == nil {
		return nil, err
	}

	return obj.(*examplev1.TestType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTestTypes) UpdateStatus(testType *examplev1.TestType) (*examplev1.TestType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithMultiTenancy(testtypesResource, "status", c.ns, testType, c.te), &examplev1.TestType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*examplev1.TestType), err
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *FakeTestTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(testtypesResource, c.ns, name, c.te), &examplev1.TestType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(testtypesResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &examplev1.TestTypeList{})
	return err
}

// Patch applies the patch and returns the patched testType.
func (c *FakeTestTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *examplev1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(testtypesResource, c.te, c.ns, name, pt, data, subresources...), &examplev1.TestType{})

	if obj == nil {
		return nil, err
	}

	return obj.(*examplev1.TestType), err
}
