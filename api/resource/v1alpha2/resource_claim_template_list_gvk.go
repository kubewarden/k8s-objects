// Code generated by GroupVersionResource generator for getting GVK data. DO NOT EDIT.

package v1alpha2

import "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime/schema"

func (v *ResourceClaimTemplateList) GroupVersionKind() schema.GroupVersionKind {
    kind := v.Kind
    apiVersion := v.APIVersion
    if kind == "" {
        kind = "ResourceClaimTemplateList"
    }
    if apiVersion == "" {
        apiVersion = SchemeGroupVersion.String()
    }

    return schema.FromAPIVersionAndKind(apiVersion, kind)
}