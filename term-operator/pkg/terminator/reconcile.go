package terminator

import (
	//"reflect"
	"github.com/sirupsen/logrus"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/labels"
)

// Reconcile new terminator CR
func Reconcile(terminator *v1alpha1.Terminator, event sdk.Event) error {
	logrus.Info(terminator)

	if event.Deleted {
		// If event is delete ignore it
		return nil
	}

	if terminator.Spec.Memcache {
		logrus.Infof("Memcache being created")
		go SendWebsocketStatus(terminator)
	}

	if terminator.Spec.Redis {
		logrus.Infof("Redis being created")
	}

	return nil
}
