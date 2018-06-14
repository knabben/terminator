package stub

import (
	"net/url"
	"context"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/gorilla/websocket"
)

func NewHandler() sdk.Handler {
	u := url.URL{
		Scheme: "ws",
		Host: "192.168.99.1:8000",
		Path: "ws/events/",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	return &Handler{conn: conn}
}

type Handler struct {
	conn *websocket.Conn
}


func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Terminator:
		err := sdk.Create(newbusyBoxPod(o))

		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create busybox pod : %v", err)
			return err
		}

		err = h.conn.WriteMessage(websocket.TextMessage,
			[]byte(string(o.Spec.Memcache)))

		if err != nil {
			logrus.Errorf("write:", err)
			return err
		}
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha1.Terminator) *v1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Terminator",
				}),
			},
			Labels: labels,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
