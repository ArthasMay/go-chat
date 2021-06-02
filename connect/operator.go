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