package connect

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gochat/proto"
	"gochat/tools"
	"time"
)

type Server struct {
	 Buckets   []*Bucket
	 Options   ServerOptions
	 bucketIdx uint32
	 operator  Operator
}

type ServerOptions struct {
	 WriteWait       time.Duration
	 PongWait        time.Duration
	 PingPeriod      time.Duration
	 MaxMessageSize  int64
	 ReadBufferSize  int
	 WriteBufferSize int
	 BroadcastSize   int
}

func NewServer(b []*Bucket, o Operator, options ServerOptions) *Server {
	s := new(Server)
	s.Buckets = b
	s.Options = options
	s.bucketIdx = uint32(len(b))
	s.operator = o
	return s
}

func (s *Server) Bucket(userId int) *Bucket {
	userIdStr := fmt.Sprintf("%d", userId)
	idx := tools.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.bucketIdx
	return s.Buckets[idx]
}

func (s *Server) writePump(ch *Channel, c *Connect) {
	// PingPeriod default eq 54s
	ticker := time.NewTicker(s.Options.PingPeriod)
	defer func() {
		ticker.Stop()
		ch.conn.Close()
	}()

	for {
		select {
		case message, ok := <- ch.broadcast:
			// write data dead time, like http timeout, default 10s
			ch.conn.SetReadDeadline(time.Now().Add(s.Options.WriteWait))
			if !ok {
				logrus.Warn("SetWriteDeadline not ok")
				ch.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := ch.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logrus.Warn("ch.conn.NextWriter err: %s", err.Error())
				return
			}
			logrus.Infof("message write body:%s", message.Body)
			w.Write(message.Body)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			ch.conn.SetWriteDeadline(time.Now().Add(s.Options.WriteWait))
			logrus.Infof("websocket.PingMessage: %v", websocket.PingMessage)
			if err := ch.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (s *Server) ReadPump(ch *Channel, c *Connect) {
	defer func() {
		logrus.Infof("start exec disconnect ...")
		if ch.Room == nil || ch.userId == 0 {
			logrus.Infof("roomId or userId eq 0")
			ch.conn.Close()
			return
		}
		logrus.Infof("exec disconnect ...")
		disconnectRequest := new(proto.DisConnectRequest)
		disconnectRequest.RoomId = ch.Room.Id
		disconnectRequest.UserId = ch.userId
		s.Bucket(ch.userId).DeleteChannel(ch)

		ch.conn.Close()
	}()

	ch.conn.SetReadLimit(s.Options.MaxMessageSize)
	ch.conn.SetReadDeadline(time.Now().Add(s.Options.PongWait))

	for {
		_, message, err := ch.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("readPump ReadMessage err:%s", err.Error())
				return
			}

			if message == nil {
				return
			}

			var connectReq *proto.ConnectRequest
			logrus.Infof("get a message :%s", message)
	        if err := json.Unmarshal([]byte(message), &connectReq); err != nil {
				logrus.Errorf("message struct %+v", connectReq)
			}
			if connectReq.AuthToken == "" {
				logrus.Errorf("s.operator.Connect no authToken")
				return
			}
			connectReq.ServerId = c.ServerId

		}
	}
}



