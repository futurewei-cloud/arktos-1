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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/kubernetes/globalscheduler/pkg/apis/dispatcher/v1"
)

// DispatcherLister helps list Dispatchers.
type DispatcherLister interface {
	// List lists all Dispatchers in the indexer.
	List(selector labels.Selector) (ret []*v1.Dispatcher, err error)
	// Dispatchers returns an object that can list and get Dispatchers.
	Dispatchers(namespace string) DispatcherNamespaceLister
	DispatchersWithMultiTenancy(namespace string, tenant string) DispatcherNamespaceLister
	DispatcherListerExpansion
}

// dispatcherLister implements the DispatcherLister interface.
type dispatcherLister struct {
	indexer cache.Indexer
}

// NewDispatcherLister returns a new DispatcherLister.
func NewDispatcherLister(indexer cache.Indexer) DispatcherLister {
	return &dispatcherLister{indexer: indexer}
}

// List lists all Dispatchers in the indexer.
func (s *dispatcherLister) List(selector labels.Selector) (ret []*v1.Dispatcher, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Dispatcher))
	})
	return ret, err
}

// Dispatchers returns an object that can list and get Dispatchers.
func (s *dispatcherLister) Dispatchers(namespace string) DispatcherNamespaceLister {
	return dispatcherNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "system"}
}

func (s *dispatcherLister) DispatchersWithMultiTenancy(namespace string, tenant string) DispatcherNamespaceLister {
	return dispatcherNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// DispatcherNamespaceLister helps list and get Dispatchers.
type DispatcherNamespaceLister interface {
	// List lists all Dispatchers in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.Dispatcher, err error)
	// Get retrieves the Dispatcher from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.Dispatcher, error)
	DispatcherNamespaceListerExpansion
}

// dispatcherNamespaceLister implements the DispatcherNamespaceLister
// interface.
type dispatcherNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all Dispatchers in the indexer for a given namespace.
func (s dispatcherNamespaceLister) List(selector labels.Selector) (ret []*v1.Dispatcher, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Dispatcher))
	})
	return ret, err
}

// Get retrieves the Dispatcher from the indexer for a given namespace and name.
func (s dispatcherNamespaceLister) Get(name string) (*v1.Dispatcher, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "system" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("dispatcher"), name)
	}
	return obj.(*v1.Dispatcher), nil
}