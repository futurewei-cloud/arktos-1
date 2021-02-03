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
	v1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LeaseLister helps list Leases.
type LeaseLister interface {
	// List lists all Leases in the indexer.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Leases returns an object that can list and get Leases.
	Leases(namespace string) LeaseNamespaceLister
	LeasesWithMultiTenancy(namespace string, tenant string) LeaseNamespaceLister
	LeaseListerExpansion
}

// leaseLister implements the LeaseLister interface.
type leaseLister struct {
	indexer cache.Indexer
}

// NewLeaseLister returns a new LeaseLister.
func NewLeaseLister(indexer cache.Indexer) LeaseLister {
	return &leaseLister{indexer: indexer}
}

// List lists all Leases in the indexer.
func (s *leaseLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Lease))
	})
	return ret, err
}

// Leases returns an object that can list and get Leases.
func (s *leaseLister) Leases(namespace string) LeaseNamespaceLister {
	return leaseNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "system"}
}

func (s *leaseLister) LeasesWithMultiTenancy(namespace string, tenant string) LeaseNamespaceLister {
	return leaseNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// LeaseNamespaceLister helps list and get Leases.
type LeaseNamespaceLister interface {
	// List lists all Leases in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Get retrieves the Lease from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.Lease, error)
	LeaseNamespaceListerExpansion
}

// leaseNamespaceLister implements the LeaseNamespaceLister
// interface.
type leaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all Leases in the indexer for a given namespace.
func (s leaseNamespaceLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Lease))
	})
	return ret, err
}

// Get retrieves the Lease from the indexer for a given namespace and name.
func (s leaseNamespaceLister) Get(name string) (*v1.Lease, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "system" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("lease"), name)
	}
	return obj.(*v1.Lease), nil
}
