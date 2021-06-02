package connect

import (
	"github.com/google/uuid"
	"fmt"
	"gochat/config"
	"runtime"
	"time"
)

var DefaultServer *Server

type Connect struct {
	ServerId string  // Connect的唯一标志id
}

func New() *Connect {
	return new(Connect)
}

// Run 启动Websocket连接服务
func (c *Connect) Run() {
	// get connect layer config
	connectConfig := config.Conf.Connect

	// set the maximum number of CPUs that can be executing
	runtime.GOMAXPROCS(connectConfig.ConnectBucket.CpuNum)

	// init logic layer rpc client, call logic layer rpc server

	// init connect layer work node: 创建Connect层: server -> buckets -> rooms/pairs -> channel
	Buckets := make([]*Bucket, connectConfig.ConnectBucket.CpuNum)
	for i := 0; i < connectConfig.ConnectBucket.CpuNum; i++ {
		Buckets[i] = NewBucket(BucketOptions{
			ChannelSize:   connectConfig.ConnectBucket.Channel,
			RoomSize:      connectConfig.ConnectBucket.Room,
			RoutineAmount: connectConfig.ConnectBucket.RoutineAmount,
			RoutineSize:   connectConfig.ConnectBucket.RoutineSize,
		})
	}
	operator := new(DefaultOperator)
	DefaultServer = NewServer(Buckets, operator, ServerOptions{
		WriteWait:       10 * time.Second,
		PongWait:        60 * time.Second,
		PingPeriod:      54 * time.Second,
		MaxMessageSize:  512,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		BroadcastSize:   512,
	})
	c.ServerId = fmt.Sprintf("%s-%s", "ws", uuid.New().String())

	// init Connect layer rpc server ,task layer will call this

	// start Connect layer server handler persistent connection

}