// Code generated by GroupVersionResource generator for getting GVK data. DO NOT EDIT.

package v2alpha1

import "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime/schema"

func (v *CronJob) GroupVersionKind() schema.GroupVersionKind {
    kind := v.Kind
    apiVersion := v.APIVersion
    if kind == "" {
        kind = "CronJob"
    }
    if apiVersion == "" {
        apiVersion = SchemeGroupVersion.String()
    }

    return schema.FromAPIVersionAndKind(apiVersion, kind)
}
