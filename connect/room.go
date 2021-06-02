package connect

import (
	"errors"
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

func (r *Room) Put(ch *Channel) (err error) {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	if !r.drop {
		if r.next != nil {
			r.next.Prev = ch
		}
		ch.Next = r.next
		ch.Prev = nil
		r.next = ch
		r.OnlineCount ++
	} else {
		err = errors.New("room drop")
	}
	return
}

func (r *Room) push(msg *proto.Msg) {
	r.rLock.RLock()
	for ch := r.next; ch != nil; ch = ch.Next {
		ch.Push(msg)
	}
    r.rLock.RUnlock()
	return
}

func (r *Room) DeleteChannel(ch *Channel) bool {
	// 为什么这里是读锁
	r.rLock.RLock()
	if ch.Next != nil { // if not footer
		ch.Next.Prev = ch.Prev
	}
	if ch.Prev != nil { // if not header
		ch.Prev.Next = ch.Next
	} else {
		r.next = ch.Next
	}

	r.OnlineCount --
	r.drop = false
	if r.OnlineCount <= 0 {
		r.drop = true
	}
	r.rLock.RUnlock()
	return r.drop
}