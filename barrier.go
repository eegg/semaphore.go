package barrier

import (
	"./semaphore"
	"./counter"
)


type Barrier struct {
	counter counter.Counter
	barrier semaphore.Semaphore
	end semaphore.Semaphore
	numThreads uint32
	funcs []func(*Barrier)
}


func (this *Barrier) Wait() {
	this.counter.IncAnd(this.numThreads, func() {
		this.barrier.V()
	})
	
	this.barrier.Turnstile()
}

func (this *Barrier) End() {
	this.counter.DecAnd(0, func() {
		this.end.V()
	})
}


func NewBarrier (funcs []func(*Barrier)) *Barrier {
	return &Barrier {
	counter: counter.Counter { Mutex: 1 },
	numThreads: uint32(len(funcs)),
	funcs: funcs,
	}
}


func (this *Barrier) Run() (sem *semaphore.Semaphore) {
	for _, f := range this.funcs {	go f(this) }
	return &this.end // caller waits on this
}