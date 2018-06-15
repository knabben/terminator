package terminator

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

//SendWebsocketStatus
func SendWebsocketStatus(term *v1alpha1.Terminator) error {
	conn, err := ConnectWebsocket()

	if err != nil {
		logrus.Warn(err)
		return err
	}

	data, err := json.Marshal(term)
	return conn.WriteMessage(websocket.TextMessage, data)
}
