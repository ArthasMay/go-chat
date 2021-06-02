package connect

import (
	"gochat/config"
	"net/http"
)

func (c *Connect) InitWebsocket() error {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {

	})
	err := http.ListenAndServe(config.Conf.Connect.ConnectWebsocket.Bind, nil)
	return err
}

func (c *Connect) serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	// after handshake, upgrade websocket
}