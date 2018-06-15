package terminator

import (
	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// podList returns a v1.PodList object
func podList() *v1.PodList {
	return &v1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}

}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []v1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)

	}
	return podNames

}

// asOwner returns an OwnerReference set as the terminator CR
func asOwner(m *v1alpha1.Terminator) metav1.OwnerReference {
	trueVar := true
	return metav1.OwnerReference{
		APIVersion: m.APIVersion,
		Kind:       m.Kind,
		Name:       m.Name,
		UID:        m.UID,
		Controller: &trueVar,
	}

}

// labelsFor returns the labels for selecting the resources
// belonging to the given terminator CR name.
func labelsFor(name, termType string) map[string]string {
	return map[string]string{
		"app": name, "term-type": termType, "hasta": "la-vista"}

}

// addOwnerRefToObject appends the desired OwnerReference to the object
func addOwnerRefToObject(obj metav1.Object, ownerRef metav1.OwnerReference) {
	obj.SetOwnerReferences(append(obj.GetOwnerReferences(), ownerRef))

}
