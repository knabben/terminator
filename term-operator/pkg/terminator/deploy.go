package terminator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

// deploymentForRabbitmq returns a memcached Deployment object
func deploymentForRabbit(term *v1alpha1.Terminator) *appsv1.Deployment {
	name := fmt.Sprintf("%s-%s", term.Name, "rabbitmq")
	selectors := labelsFor(term.Name, "rabbitmq")

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: selectors,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectors,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image: "bitnami/rabbitmq:3.7",
						Name:  "rabbitmq",
						Ports: []v1.ContainerPort{{
							ContainerPort: 5672,
							Name:          "rabbitmq",
						}},
						Env: []v1.EnvVar{
							{
								Name:  "RABBITMQ_USERNAME",
								Value: "guest",
							},
							{
								Name:  "RABBITMQ_PASSWORD",
								Value: "guest",
							},
						},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(term))
	err := sdk.Create(dep)

	podNames := getPodList(selectors, term.Namespace)
	if !reflect.DeepEqual(podNames, term.Status.RabbitmqNode) {
		term.Status.RabbitmqNode = podNames
	}
	setOperatorStatus(term)

	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Error(err)
	}

	svc := &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.GetNamespace(),
			Labels:    selectors,
		},
		Spec: v1.ServiceSpec{
			Selector: selectors,
			Ports: []v1.ServicePort{
				{
					Name:     "rabbitmq",
					Protocol: v1.ProtocolTCP,
					Port:     5672,
				},
			},
		},
	}
	addOwnerRefToObject(svc, asOwner(term))
	err = sdk.Create(svc)
	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Error("failed to create memcache service: %v", err)
	}

	return dep
}

// deploymentForMemcached returns a memcached Deployment object
func deploymentForMemcached(term *v1alpha1.Terminator) *appsv1.Deployment {
	name := fmt.Sprintf("%s-%s", term.Name, "memcache")
	selectors := labelsFor(term.Name, "memcache")

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: selectors,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectors,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image:   "memcached:1.5.6-alpine",
						Name:    "memcached",
						Command: []string{"memcached", "-o", "modern", "-v"},
						Ports: []v1.ContainerPort{{
							ContainerPort: 11211,
							Name:          "memcached",
						}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(term))
	err := sdk.Create(dep)

	podNames := getPodList(selectors, term.Namespace)
	if !reflect.DeepEqual(podNames, term.Status.MemcacheNode) {
		term.Status.MemcacheNode = podNames
	}
	setOperatorStatus(term)

	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Error(err)
	}

	svc := &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.GetNamespace(),
			Labels:    selectors,
		},
		Spec: v1.ServiceSpec{
			Selector: selectors,
			Ports: []v1.ServicePort{
				{
					Name:     "memcache",
					Protocol: v1.ProtocolTCP,
					Port:     11211,
				},
			},
		},
	}
	addOwnerRefToObject(svc, asOwner(term))
	err = sdk.Create(svc)
	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Error("failed to create memcache service: %v", err)
	}

	return dep
}

// deploymentForRedis creates a redis Deployment object
func deploymentForRedis(term *v1alpha1.Terminator) *appsv1.Deployment {
	name := fmt.Sprintf("%s-%s", term.Name, "redis")
	selectors := labelsFor(term.Name, "redis")

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: selectors,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectors,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image: "bitnami/redis:4.0.10",
						Name:  "redis",
						Env: []v1.EnvVar{{
							Name:  "ALLOW_EMPTY_PASSWORD",
							Value: "yes",
						}},
						Ports: []v1.ContainerPort{{
							ContainerPort: 6379,
							Name:          "redis",
						}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(term))
	err := sdk.Create(dep)

	podNames := getPodList(selectors, term.Namespace)
	if !reflect.DeepEqual(podNames, term.Status.RedisNode) {
		term.Status.RedisNode = podNames
	}
	setOperatorStatus(term)

	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Info(err)
	}

	svc := &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: term.GetNamespace(),
			Labels:    selectors,
		},
		Spec: v1.ServiceSpec{
			Selector: selectors,
			Ports: []v1.ServicePort{
				{
					Name:     "redis",
					Protocol: v1.ProtocolTCP,
					Port:     6379,
				},
			},
		},
	}
	addOwnerRefToObject(svc, asOwner(term))
	err = sdk.Create(svc)
	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Error("failed to create redis service: %v", err)
	}

	return dep
}

func setReplica(obj *appsv1.Deployment, replicas int32) error {
	err := sdk.Get(obj)
	if err != nil {
		return err
	}

	if *obj.Spec.Replicas != replicas {
		obj.Spec.Replicas = &replicas
		err = sdk.Update(obj)
		if err != nil {
			return err
		}
		logrus.Warn("Setting replica to: ", replicas)
	}
	return nil
}

func setOperatorStatus(term *v1alpha1.Terminator) error {
	err := sdk.Update(term)
	if err != nil {
		logrus.Errorf("failed to update memcached status: %v", err)
		return err

	}

	return nil
}
