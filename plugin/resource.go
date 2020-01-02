package main

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// crdBackingService returns the GVR and resources
func crdBackingService() (schema.GroupVersionResource, *unstructured.Unstructured) {
	gvr := schema.GroupVersionResource{
		Group:    "backing.bluebird.io",
		Version:  "v1",
		Resource: "services",
	}
	resource := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "backing.bluebird.io/v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name": "name-1",
			},
			"spec": map[string]interface{}{
				"name": "redis",
			},
		},
	}
	return gvr, resource
}


