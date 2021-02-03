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

package v1beta1

import (
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionLister helps list CustomResourceDefinitions.
type CustomResourceDefinitionLister interface {
	// List lists all CustomResourceDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error)
	// CustomResourceDefinitions returns an object that can list and get CustomResourceDefinitions.
	CustomResourceDefinitions() CustomResourceDefinitionTenantLister
	CustomResourceDefinitionsWithMultiTenancy(tenant string) CustomResourceDefinitionTenantLister
	// Get retrieves the CustomResourceDefinition from the index for a given name.
	Get(name string) (*v1beta1.CustomResourceDefinition, error)
	CustomResourceDefinitionListerExpansion
}

// customResourceDefinitionLister implements the CustomResourceDefinitionLister interface.
type customResourceDefinitionLister struct {
	indexer cache.Indexer
}

// NewCustomResourceDefinitionLister returns a new CustomResourceDefinitionLister.
func NewCustomResourceDefinitionLister(indexer cache.Indexer) CustomResourceDefinitionLister {
	return &customResourceDefinitionLister{indexer: indexer}
}

// List lists all CustomResourceDefinitions in the indexer.
func (s *customResourceDefinitionLister) List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CustomResourceDefinition))
	})
	return ret, err
}

// Get retrieves the CustomResourceDefinition from the index for a given name.
func (s *customResourceDefinitionLister) Get(name string) (*v1beta1.CustomResourceDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("customresourcedefinition"), name)
	}
	return obj.(*v1beta1.CustomResourceDefinition), nil
}

// CustomResourceDefinitions returns an object that can list and get CustomResourceDefinitions.
func (s *customResourceDefinitionLister) CustomResourceDefinitions() CustomResourceDefinitionTenantLister {
	return customResourceDefinitionTenantLister{indexer: s.indexer, tenant: "system"}
}

func (s *customResourceDefinitionLister) CustomResourceDefinitionsWithMultiTenancy(tenant string) CustomResourceDefinitionTenantLister {
	return customResourceDefinitionTenantLister{indexer: s.indexer, tenant: tenant}
}

// CustomResourceDefinitionTenantLister helps list and get CustomResourceDefinitions.
type CustomResourceDefinitionTenantLister interface {
	// List lists all CustomResourceDefinitions in the indexer for a given tenant/tenant.
	List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error)
	// Get retrieves the CustomResourceDefinition from the indexer for a given tenant/tenant and name.
	Get(name string) (*v1beta1.CustomResourceDefinition, error)
	CustomResourceDefinitionTenantListerExpansion
}

// customResourceDefinitionTenantLister implements the CustomResourceDefinitionTenantLister
// interface.
type customResourceDefinitionTenantLister struct {
	indexer cache.Indexer
	tenant  string
}

// List lists all CustomResourceDefinitions in the indexer for a given tenant.
func (s customResourceDefinitionTenantLister) List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error) {
	err = cache.ListAllByTenant(s.indexer, s.tenant, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CustomResourceDefinition))
	})
	return ret, err
}

// Get retrieves the CustomResourceDefinition from the indexer for a given tenant and name.
func (s customResourceDefinitionTenantLister) Get(name string) (*v1beta1.CustomResourceDefinition, error) {
	key := s.tenant + "/" + name
	if s.tenant == "system" {
		key = name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("customresourcedefinition"), name)
	}
	return obj.(*v1beta1.CustomResourceDefinition), nil
}
