package proto

type Msg struct {
	Ver       int `json:"ver"`  // protocol version
	Operation int `json:"op"`
	SeqId     string `json:"seq"`
	Body      []byte `json:"body"`
}

type PushRoomMsgRequest struct {
	RoomId int
	Msg Msg
}


