package twophasebarrier

import (
	"./semaphore"
	"./counter"
)


type TwoPhaseBarrier struct {
	counter counter.Counter
	endCounter counter.Counter
	barrierOne semaphore.Semaphore
	barrierTwo semaphore.Semaphore
	end semaphore.Semaphore
	numThreads uint32
	funcs []func(*TwoPhaseBarrier)
}


func (this *TwoPhaseBarrier) StartOfLoop() {
	this.counter.IncAnd(this.numThreads, func() {
			this.barrierTwo.P()
			this.barrierOne.V()
	})

	this.barrierOne.Turnstile()
}

func (this *TwoPhaseBarrier) EndOfLoop() {
	this.counter.DecAnd(0, func() {
			this.barrierOne.P()
			this.barrierTwo.V()
	})

	this.barrierTwo.Turnstile()	
}

func (this *TwoPhaseBarrier) End() {
	this.endCounter.IncAnd(this.numThreads, func() {
		this.end.V()
	})
}


func NewTwoPhaseBarrier (funcs []func(*TwoPhaseBarrier)) *TwoPhaseBarrier {
	return &TwoPhaseBarrier {
	counter: counter.Counter { Mutex: 1 },
	endCounter: counter.Counter { Mutex: 1 },
	barrierTwo: semaphore.Semaphore(1),
	numThreads: uint32(len(funcs)),
	funcs: funcs,
	}
}


func (this *TwoPhaseBarrier) Run() (sem *semaphore.Semaphore) {
	for _, f := range this.funcs {	go f(this) }
	return &this.end // caller waits on this
}