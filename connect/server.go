package connect

import "time"

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

//func (s *Server) Bucket(userId int) *Bucket {
//
//}

