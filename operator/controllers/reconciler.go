package controllers

/* Work based on: https://github.com/DirectXMan12/kubebuilder-workshops/blob/kubecon-us-2019/controllers/helpers.go */

import (
	backingv1 "github.com/knabben/terminator/operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"strconv"
)

func (r *ServiceReconciler) ServiceDeployment(service backingv1.Service) (appsv1.Deployment, error) {
	defOne := int32(1)

	depl := appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: appsv1.SchemeGroupVersion.String(),
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      service.Name,
			Namespace: service.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &defOne,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"redis": service.Name, "role": "leader"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"redis": service.Name, "role": "leader"},
				},
				Spec: corev1.PodSpec{
					Containers: r.generateContainers(),
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(&service, &depl, r.Scheme); err != nil {
		return depl, err
	}

	return depl, nil
}

func (r *ServiceReconciler) generateContainers() []corev1.Container {
	return []corev1.Container{
		{
			Name:  "redis",
			Image: "k8s.gcr.io/redis:e2e",
			Ports: []corev1.ContainerPort{
				{ContainerPort: 6379, Name: "redis", Protocol: "TCP"},
			},
		},
	}
}

func (r *ServiceReconciler) ServiceService(service backingv1.Service) (corev1.Service, error) {
	port, _ := strconv.Atoi(service.Spec.Port)

	svc := corev1.Service{
		TypeMeta: metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Service"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      service.Name,
			Namespace: service.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Name: "host", Port: int32(port), Protocol: "TCP", TargetPort: intstr.FromInt(port)},
			},
			Selector: map[string]string{"role": "leader"},
			Type:     corev1.ServiceTypeLoadBalancer,
		},
	}

	// always set the controller reference so that we know which object owns this.
	if err := ctrl.SetControllerReference(&service, &svc, r.Scheme); err != nil {
		return svc, err
	}

	return svc, nil
}
