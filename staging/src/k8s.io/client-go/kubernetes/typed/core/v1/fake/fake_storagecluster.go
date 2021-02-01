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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageClusters implements StorageClusterInterface
type FakeStorageClusters struct {
	Fake *FakeCoreV1
}

var storageclustersResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "storageclusters"}

var storageclustersKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "StorageCluster"}

// Get takes name of the storageCluster, and returns the corresponding storageCluster object, and an error if there is any.
func (c *FakeStorageClusters) Get(name string, options v1.GetOptions) (result *corev1.StorageCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storageclustersResource, name), &corev1.StorageCluster{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.StorageCluster), err
}

// List takes label and field selectors, and returns the list of StorageClusters that match those selectors.
func (c *FakeStorageClusters) List(opts v1.ListOptions) (result *corev1.StorageClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storageclustersResource, storageclustersKind, opts), &corev1.StorageClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.StorageClusterList{ListMeta: obj.(*corev1.StorageClusterList).ListMeta}
	for _, item := range obj.(*corev1.StorageClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested storageClusters.
func (c *FakeStorageClusters) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storageclustersResource, opts))
	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a storageCluster and creates it.  Returns the server's representation of the storageCluster, and an error, if there is any.
func (c *FakeStorageClusters) Create(storageCluster *corev1.StorageCluster) (result *corev1.StorageCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storageclustersResource, storageCluster), &corev1.StorageCluster{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.StorageCluster), err
}

// Update takes the representation of a storageCluster and updates it. Returns the server's representation of the storageCluster, and an error, if there is any.
func (c *FakeStorageClusters) Update(storageCluster *corev1.StorageCluster) (result *corev1.StorageCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storageclustersResource, storageCluster), &corev1.StorageCluster{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.StorageCluster), err
}

// Delete takes name of the storageCluster and deletes it. Returns an error if one occurs.
func (c *FakeStorageClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storageclustersResource, name), &corev1.StorageCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {

	action := testing.NewRootDeleteCollectionAction(storageclustersResource, listOptions)
	_, err := c.Fake.Invokes(action, &corev1.StorageClusterList{})
	return err
}

// Patch applies the patch and returns the patched storageCluster.
func (c *FakeStorageClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.StorageCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageclustersResource, name, pt, data, subresources...), &corev1.StorageCluster{})
	if obj == nil {
		return nil, err
	}

	return obj.(*corev1.StorageCluster), err
}
