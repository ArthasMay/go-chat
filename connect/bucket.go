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

func (b *Bucket) Put(userId int, roomId int, ch *Channel) (err error) {
	var (
		room *Room
		ok bool
	)
	b.cLock.Lock()
	if roomId != NoRoom {
		if room, ok = b.rooms[roomId]; !ok {
			room = NewRoom(roomId)
			b.rooms[roomId] = room
		}
		ch.Room = room
	}
	b.chs[userId] = ch
	ch.userId = userId
	b.cLock.Unlock()

	if room != nil {
		room.Put(ch)
	}
 	return
}

func (b *Bucket) DeleteChannel(ch *Channel) {
	var (
		ok bool
		room *Room
	)

	b.cLock.RLock()
	if ch, ok = b.chs[ch.userId]; ok {
		room = ch.Room
		delete(b.chs, ch.userId)
	}
	if room != nil {
		if room.DeleteChannel(ch) && room.drop == true {
			delete(b.rooms, room.Id)
		}
	}
	b.cLock.RUnlock()
}

func (b *Bucket) Channel(userId int) (ch *Channel) {
	b.cLock.RLock()
	ch = b.chs[userId]
	b.cLock.RUnlock()
	return
}

func (b *Bucket) BroadcastRoom(pushRoomMsgReq *proto.PushRoomMsgRequest) {

}