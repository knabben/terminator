package stub

import (
	"encoding/json"
	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// ConnectWebSocket start a long poll connection
func ConnectWebsocket() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "192.168.99.1:8000", Path: "ws/events/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

func (h *Handler) tryConnect() error {
	conn, err := ConnectWebsocket()
	if err != nil {
		logrus.Warn(err)
		return err
	}
	h.conn = conn
	return nil
}

//SendWebsocketStatus
func (h *Handler) SendWebsocketStatus(term *v1alpha1.Terminator) error {
	data, _ := json.Marshal(term)

	if h.conn == nil {
		h.tryConnect()
	}

	err := h.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		h.tryConnect()
		return err
	}
	return nil
}
