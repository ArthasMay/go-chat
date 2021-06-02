package connect

import (
	"gochat/proto"
	"sync"
)

type Bucket struct {
	cLock         sync.RWMutex
	chs           map[int]*Channel
	bucketOptions BucketOptions
	rooms         map[int]*Room
	routines      []chan *proto.PushRoomMsgRequest
	routinesNum   int
	broadcast     chan []byte
}

type BucketOptions struct {
	 ChannelSize int
	 RoomSize int
	 RoutineAmount uint64
	 RoutineSize int
}

func NewBucket(bucketOptions BucketOptions) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[int]*Channel, bucketOptions.ChannelSize)
	b.bucketOptions = bucketOptions
	b.routines = make([]chan *proto.PushRoomMsgRequest, bucketOptions.RoutineAmount)
	b.rooms = make(map[int]*Room, bucketOptions.RoomSize)

	for i := uint64(0); i < b.bucketOptions.RoutineAmount; i++ {
		c := make(chan *proto.PushRoomMsgRequest, b.bucketOptions.RoutineSize)
		b.routines[i] = c
		go b.PushRoom(c)
	}
	return
}

func (b *Bucket) PushRoom(ch chan *proto.PushRoomMsgRequest) {
	for {
		var (
			arg *proto.PushRoomMsgRequest
			room *Room
		)
		arg = <-ch
		if room = b.Room(arg.RoomId); room != nil {
			room.push(&arg.Msg)
		}
	}
}

func (b *Bucket) Room(rid int) (room *Room) {
	b.cLock.RLock()
	room, _ = b.rooms[rid]
	b.cLock.RUnlock()
	return
}