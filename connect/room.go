package connect

import (
	"gochat/proto"
	"sync"
)

const NoRoom = -1

type Room struct {
	Id          int
	OnlineCount int
	rLock       sync.RWMutex
	drop        bool
	next        *Channel
}

func NewRoom(roomId int) *Room {
	room := new(Room)
	room.Id = roomId
	room.drop = false
	room.OnlineCount = 0
	room.next = nil
	return room
}

func (r *Room) push(msg *proto.Msg) {
	r.rLock.RLock()
	for ch := r.next; ch != nil; ch = ch.Next {
		ch.Push(msg)
	}
    r.rLock.RUnlock()
	return
}