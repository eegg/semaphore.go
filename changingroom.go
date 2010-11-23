package changingroom

import (
	"fmt"
	"./semaphore"
	"./counter"
	)

type ChangingRoom struct {
	manQueue counter.Counter
	womenInRoom counter.Counter
	changingRoom semaphore.Semaphore
}

func NewChangingRoom() *ChangingRoom {
	return &ChangingRoom {
	manQueue: counter.NewCounter(),
	womenInRoom: counter.NewCounter(),
	changingRoom: semaphore.Semaphore(0),
	}
}

func (this *ChangingRoom) ManIn() {
	this.womenInRoom.Mutex.P()
	if this.womenInRoom.Value > 0 {
		this.womenInRoom.Mutex.V()
		this.manQueue.Inc()
		this.changingRoom.P()
	} else {
		this.womenInRoom.Mutex.V()
	}
}

func (this *ChangingRoom) WomanIn() {
	this.womenInRoom.Inc()
}

func (this *ChangingRoom) WomanOut() {
	this.womenInRoom.DecAnd(0, func() {
		this.manQueue.DecEachV(&this.changingRoom)
	})
}

func (this *ChangingRoom) String() string {
	return fmt.Sprintf("men in queue: %v, women in room: %v", this.manQueue.Value, this.womenInRoom.Value)
}