package proto

type Msg struct {
	Ver       int `json:"ver"`  // protocol version
	Operation int `json:"op"`
	SeqId     string `json:"seq"`
	Body      []byte `json:"body"`
}

// PushMsgRequest p2p聊天发送的消息
type PushMsgRequest struct {
	UserId int
	Msg Msg
}

// PushRoomMsgRequest 向Room广播
type PushRoomMsgRequest struct {
	RoomId int
	Msg Msg
}


