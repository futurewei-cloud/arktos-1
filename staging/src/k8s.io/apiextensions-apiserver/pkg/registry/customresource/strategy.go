/*
Copyright 2017 The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

package customresource

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-openapi/validate"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apiserverstorage "k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	utilfeature "k8s.io/apiserver/pkg/util/feature"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	schemaobjectmeta "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/objectmeta"
	apiextensionsfeatures "k8s.io/apiextensions-apiserver/pkg/features"
)

// customResourceStrategy implements behavior for CustomResources.
type customResourceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator

	namespaceScoped bool
	validator       customResourceValidator
	schemas         map[string]*structuralschema.Structural
	status          *apiextensions.CustomResourceSubresourceStatus
	scale           *apiextensions.CustomResourceSubresourceScale
	tenantScoped    bool
}

func NewStrategy(typer runtime.ObjectTyper, tenantScoped, namespaceScoped bool, kind schema.GroupVersionKind, schemaValidator, statusSchemaValidator *validate.SchemaValidator, schemas map[string]*structuralschema.Structural, status *apiextensions.CustomResourceSubresourceStatus, scale *apiextensions.CustomResourceSubresourceScale) customResourceStrategy {
	return customResourceStrategy{
		ObjectTyper:     typer,
		NameGenerator:   names.SimpleNameGenerator,
		tenantScoped:    tenantScoped,
		namespaceScoped: namespaceScoped,
		status:          status,
		scale:           scale,
		validator: customResourceValidator{
			tenantScoped:          tenantScoped,
			namespaceScoped:       namespaceScoped,
			kind:                  kind,
			schemaValidator:       schemaValidator,
			statusSchemaValidator: statusSchemaValidator,
		},
		schemas: schemas,
	}
}

func (a customResourceStrategy) NamespaceScoped() bool {
	return a.namespaceScoped
}

func (a customResourceStrategy) TenantScoped() bool {
	return a.tenantScoped
}

// PrepareForCreate clears the status of a CustomResource before creation.
func (a customResourceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	if utilfeature.DefaultFeatureGate.Enabled(apiextensionsfeatures.CustomResourceSubresources) && a.status != nil {
		customResourceObject := obj.(*unstructured.Unstructured)
		customResource := customResourceObject.UnstructuredContent()

		// create cannot set status
		if _, ok := customResource["status"]; ok {
			delete(customResource, "status")
		}
	}

	accessor, _ := meta.Accessor(obj)
	accessor.SetGeneration(1)
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (a customResourceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newCustomResourceObject := obj.(*unstructured.Unstructured)
	oldCustomResourceObject := old.(*unstructured.Unstructured)

	newCustomResource := newCustomResourceObject.UnstructuredContent()
	oldCustomResource := oldCustomResourceObject.UnstructuredContent()

	// If the /status subresource endpoint is installed, update is not allowed to set status.
	if utilfeature.DefaultFeatureGate.Enabled(apiextensionsfeatures.CustomResourceSubresources) && a.status != nil {
		_, ok1 := newCustomResource["status"]
		_, ok2 := oldCustomResource["status"]
		switch {
		case ok2:
			newCustomResource["status"] = oldCustomResource["status"]
		case ok1:
			delete(newCustomResource, "status")
		}
	}

	// except for the changes to `metadata`, any other changes
	// cause the generation to increment.
	newCopyContent := copyNonMetadata(newCustomResource)
	oldCopyContent := copyNonMetadata(oldCustomResource)
	if !apiequality.Semantic.DeepEqual(newCopyContent, oldCopyContent) {
		oldAccessor, _ := meta.Accessor(oldCustomResourceObject)
		newAccessor, _ := meta.Accessor(newCustomResourceObject)
		newAccessor.SetGeneration(oldAccessor.GetGeneration() + 1)
	}
}

func copyNonMetadata(original map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for key, val := range original {
		if key == "metadata" {
			continue
		}
		ret[key] = val
	}
	return ret
}

// Validate validates a new CustomResource.
func (a customResourceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	var errs field.ErrorList
	errs = append(errs, a.validator.Validate(ctx, obj, a.scale)...)

	// validate embedded resources
	if u, ok := obj.(*unstructured.Unstructured); ok {
		v := obj.GetObjectKind().GroupVersionKind().Version
		errs = append(errs, schemaobjectmeta.Validate(nil, u.Object, a.schemas[v], false)...)
	}

	return errs
}

// Canonicalize normalizes the object after validation.
func (customResourceStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for CustomResources; this means a POST is
// needed to create one.
func (customResourceStrategy) AllowCreateOnUpdate() bool {
	return false
}

// AllowUnconditionalUpdate is the default update policy for CustomResource objects.
func (customResourceStrategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user updating status.
func (a customResourceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	var errs field.ErrorList
	errs = append(errs, a.validator.ValidateUpdate(ctx, obj, old, a.scale)...)

	// Checks the embedded objects. We don't make a difference between update and create for those.
	if u, ok := obj.(*unstructured.Unstructured); ok {
		v := obj.GetObjectKind().GroupVersionKind().Version
		errs = append(errs, schemaobjectmeta.Validate(nil, u.Object, a.schemas[v], false)...)
	}

	return errs
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func (a customResourceStrategy) GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return nil, nil, err
	}
	crdObj := obj.(*unstructured.Unstructured)
	crdRes := crdObj.UnstructuredContent()
	fieldSet := objectMetaFieldsSet(accessor, a.namespaceScoped, a.tenantScoped)
	if statusVal, ok := crdRes["status"]; ok {
		if statusRes, ok := statusVal.(map[string]interface{}); ok {
			if phase, ok := statusRes["phase"]; ok {
				fieldSet["status.phase"] = fmt.Sprintf("%v", phase)
			}
			if distributorname, ok := statusRes["distributor_name"]; ok {
				fieldSet["status.distributor_name"] = fmt.Sprintf("%v", distributorname)
			}
			if schedulername, ok := statusRes["scheduler_name"]; ok {
				fieldSet["status.scheduler_name"] = fmt.Sprintf("%v", schedulername)
			}
			if dispatchername, ok := statusRes["dispatcher_name"]; ok {
				fieldSet["status.dispatcher_name"] = fmt.Sprintf("%v", dispatchername)
			}
			if clusternames, ok := statusRes["cluster_names"]; ok {
				if clusters, ok := clusternames.([]interface{}); ok {
					if len(clusters) > 0 {
						fieldSet["status.cluster_name"] = fmt.Sprintf("%v", clusters[0])
					}
				}
			}
		}
	}
	return labels.Set(accessor.GetLabels()), fieldSet, nil
}

// objectMetaFieldsSet returns a fields that represent the ObjectMeta.
func objectMetaFieldsSet(objectMeta metav1.Object, namespaceScoped, tenantScoped bool) fields.Set {
	if namespaceScoped {
		return fields.Set{
			"metadata.name":      objectMeta.GetName(),
			"metadata.namespace": objectMeta.GetNamespace(),
			"metadata.tenant":    objectMeta.GetTenant(),
			"metadata.hashkey":   strconv.FormatInt(objectMeta.GetHashKey(), 10),
		}
	}
	if tenantScoped {
		return fields.Set{
			"metadata.name":    objectMeta.GetName(),
			"metadata.tenant":  objectMeta.GetTenant(),
			"metadata.hashkey": strconv.FormatInt(objectMeta.GetHashKey(), 10),
		}
	}
	return fields.Set{
		"metadata.name":    objectMeta.GetName(),
		"metadata.hashkey": strconv.FormatInt(objectMeta.GetHashKey(), 10),
	}
}

// MatchCustomResourceDefinitionStorage is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func (a customResourceStrategy) MatchCustomResourceDefinitionStorage(label labels.Selector, field fields.Selector) apiserverstorage.SelectionPredicate {
	return apiserverstorage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: a.GetAttrs,
	}
}
