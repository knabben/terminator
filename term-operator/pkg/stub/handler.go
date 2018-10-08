package stub

import (
	"context"

	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/knabben/terminator/term-operator/pkg/terminator"

	"github.com/gorilla/websocket"
	"github.com/operator-framework/operator-sdk/pkg/sdk"

	"github.com/sirupsen/logrus"
)

func NewHandler() sdk.Handler {
	conn, err := ConnectWebsocket()
	if err != nil {
		logrus.Warn(err)
	}
	return &Handler{conn: conn}
}

type Handler struct {
	conn *websocket.Conn
}

// Handle starts the websocket
func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Terminator:
		terminator.Reconcile(o, event)

		// Send terminator status via websockets
		h.SendWebsocketStatus(o)

	}
	return nil
}
