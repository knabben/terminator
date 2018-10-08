package terminator

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/oleiade/reflections"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func startDeployment(term *v1alpha1.Terminator, extra map[string]string, envVar []v1.EnvVar) *appsv1.Deployment {
	slug := extra["name"]
	port, err := strconv.ParseInt(extra["port"], 10, 32)

	name := fmt.Sprintf("%s-%s", term.Name, slug)
	selectors := labelsFor(term.Name, slug)

	// Create a new deployment
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
						Image: extra["image"],
						Name:  slug,
						Env:   envVar,
						Ports: []v1.ContainerPort{{
							ContainerPort: int32(port),
							Name:          extra["name"],
						}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(term))
	err = sdk.Create(dep)
	if err != nil && !errors.IsAlreadyExists(err) {
		raven.CaptureError(err, nil)
		logrus.Error(err)
	}

	// Create a new service
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
					Name:     extra["name"],
					Port:     int32(port),
					Protocol: v1.ProtocolTCP,
				},
			},
		},
	}
	addOwnerRefToObject(svc, asOwner(term))
	err = sdk.Create(svc)
	if err != nil && !errors.IsAlreadyExists(err) {
		raven.CaptureError(err, nil)
		logrus.Error("failed to create service: %v", err)
	}

	// Set status on Terminator structure
	podNames := getPodList(selectors, term.Namespace)
	err = reflections.SetField(&term.Status, extra["status"], podNames)
	if err != nil {
		raven.CaptureError(err, nil)
		logrus.Error("Failed to set node status: %v", err)
	}
	setOperatorStatus(term)

	return dep
}

// deploymentForRabbitmq returns a memcached Deployment object
func deploymentForRabbit(term *v1alpha1.Terminator) *appsv1.Deployment {
	extra := map[string]string{
		"image":  "bitnami/rabbitmq:3.7",
		"name":   "rabbitmq",
		"port":   "5672",
		"status": "RabbitmqNode",
	}

	envVars := []v1.EnvVar{{
		Name:  "RABBITMQ_USERNAME",
		Value: "guest",
	}, {
		Name:  "RABBITMQ_PASSWORD",
		Value: "guest",
	}}
	return startDeployment(term, extra, envVars)
}

// deploymentForMemcached returns a memcached Deployment object
func deploymentForMemcached(term *v1alpha1.Terminator) *appsv1.Deployment {
	//Command: []string{"memcached", "-o", "modern", "-v"},

	extra := map[string]string{
		"image":  "memcached:1.5.6-alpine",
		"name":   "memcached",
		"port":   "11211",
		"status": "MemcacheNode",
	}
	return startDeployment(term, extra, []v1.EnvVar{})
}

// deploymentForRedis creates a redis Deployment object
func deploymentForRedis(term *v1alpha1.Terminator) *appsv1.Deployment {
	extra := map[string]string{
		"image":  "bitnami/redis:4.0.10",
		"name":   "redis",
		"port":   "6379",
		"status": "RedisNode",
	}
	envVars := []v1.EnvVar{{
		Name:  "ALLOW_EMPTY_PASSWORD",
		Value: "yes",
	}}

	return startDeployment(term, extra, envVars)
}

// deploaymentForElastic a elastic Deployment object
func deploymentForElastic(term *v1alpha1.Terminator) *appsv1.Deployment {
	extra := map[string]string{
		"image":  "elasticsearch:6.4.0",
		"name":   "elastic",
		"port":   "9200",
		"status": "ElasticNode",
	}
	return startDeployment(term, extra, []v1.EnvVar{})
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
			raven.CaptureError(err, nil)
			return err
		}
		logrus.Warn("Setting replica to: ", replicas)
	}
	return nil
}

func setOperatorStatus(term *v1alpha1.Terminator) error {
	err := sdk.Update(term)
	if err != nil {
		raven.CaptureError(err, nil)
		logrus.Errorf("failed to update status: %v", err)
		return err

	}

	return nil
}
