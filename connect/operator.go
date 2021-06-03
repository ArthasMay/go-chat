package connect

import "gochat/proto"

type Operator interface {
	//Connect(conn *proto.ConnectRequest) (int, error)
	//Disconnect(disConn *proto.DisConnectRequest) (err error)

	ConnectWithoutRPC(conn *proto.ConnectRequest) (userId int, err error)
	DisconnectWithoutRPC(disconn *proto.DisConnectRequest) (err error)
}

type DefaultOperator struct {

}

func (o *DefaultOperator)ConnectWithoutRPC(conn *proto.ConnectRequest) (userId int, err error) {
	userId = 100001
	return
}

func (o *DefaultOperator)DisconnectWithoutRPC(disconn *proto.DisConnectRequest) (err error) {
	return
}