package stub

import (
	"encoding/json"
	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"net/url"

	"github.com/gorilla/websocket"
)

// ConnectWebSocket start a long poll connection
func ConnectWebsocket() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "192.168.99.1:8000", Path: "ws/events/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return conn, err
}

//SendWebsocketStatus
func SendWebsocketStatus(term *v1alpha1.Terminator, conn *websocket.Conn) error {
	data, _ := json.Marshal(term)
	return conn.WriteMessage(websocket.TextMessage, data)
}
