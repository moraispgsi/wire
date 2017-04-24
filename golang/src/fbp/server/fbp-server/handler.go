package fbpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type request struct {
	SubProtocol string                 `json:"sub-protocol"`
	Topic       string                 `json:"topic"`
	Payload     map[string]interface{} `json:"payload"`
}

type runtimePort struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Addressable bool   `json:"addressable"`
	Required    bool   `json:"required"`
}

type runtimePorts struct {
	Graph    string      `json:"graph"`
	InPorts  runtimePort `json:"inPorts"`
	OutPorts runtimePort `json:"outPorts"`
}

type runtimePacket struct {
	Port    string                 `json:"port"`
	Event   string                 `json:"event"`
	Payload map[string]interface{} `json:"payload"`
	Graph   string                 `json:"graph"`
	Secret  string                 `json:"secret"`
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {

	var conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func(conn *websocket.Conn) {
		for {
			var req request

			err := conn.ReadJSON(&req)
			if err != nil {
				conn.Close()
				return
			}

			switch {
			case req.SubProtocol == "runtime":
				switch {
				case req.Topic == "getruntime":

					type payloadResponse struct {
						Type            string   `json:"type"`
						Version         float64  `json:"version"`
						Capabilities    []string `json:"capabilities"`
						AllCapabilities []string `json:"allCapabilities"`
						Label           string   `json:"label"`
					}
					type runtimeResponse struct {
						SubProtocol string          `json:"sub-protocol"`
						Topic       string          `json:"topic"`
						Payload     payloadResponse `json:"payload"`
					}

					json := runtimeResponse{
						SubProtocol: "runtime",
						Topic:       "runtime",
						Payload: payloadResponse{
							Type:            "gofbp",
							Version:         0.0,
							Capabilities:    []string{"protocol:runtime"},
							AllCapabilities: []string{"protocol:runtime"},
							Label:           "Golang FBP runtime",
						},
					}

					conn.WriteJSON(json)
				case req.Topic == "packet":

				}
			}
			fmt.Println("Mensagem recebida: " + req.SubProtocol)
		}
	}(conn)

}
