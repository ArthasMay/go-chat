package connect

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gochat/config"
	"net/http"
)

func (c *Connect) InitWebsocket() error {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		c.serveWs(DefaultServer, writer, request)
	})
	err := http.ListenAndServe(config.Conf.Connect.ConnectWebsocket.Bind, nil)
	return err
}

func (c *Connect) serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	// after handshake, upgrade websocket
	upgrader := websocket.Upgrader{
		ReadBufferSize: server.Options.ReadBufferSize,
		WriteBufferSize: server.Options.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if  err != nil {
		logrus.Errorf("serverWs err:%s", err.Error())
		return
	}

	var ch *Channel
	// default broadcast size eq 512
	ch = NewChannel(server.Options.BroadcastSize)
	ch.conn = conn
	// send data to websocket conn
	go server.writePump(ch, c)
    // get data from websocket conn
    go server.ReadPump(ch, c)
}