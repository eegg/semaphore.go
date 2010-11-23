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
	numThreads int
	funcs []func(*TwoPhaseBarrier)
}

func (this *TwoPhaseBarrier) OpenAndShut(
	goThrough *semaphore.Semaphore,
	stopAt *semaphore.Semaphore,
	amt int, check int) {

	this.counter.AddAnd(amt, check, func() {
		stopAt.P()
		goThrough.V()
	})

	goThrough.Turnstile()
}

func (this *TwoPhaseBarrier) StartOfLoop() {
	this.OpenAndShut(&this.barrierOne, &this.barrierTwo, 1, this.numThreads)
}

func (this *TwoPhaseBarrier) EndOfLoop() {
	this.OpenAndShut(&this.barrierTwo, &this.barrierOne, -1, 0)
}

func (this *TwoPhaseBarrier) End() {
	this.endCounter.IncAnd(this.numThreads, func() {
		this.end.V()
	})
}


func NewTwoPhaseBarrier (funcs []func(*TwoPhaseBarrier)) *TwoPhaseBarrier {
	return &TwoPhaseBarrier {
	counter: counter.NewCounter(),
	endCounter: counter.NewCounter(),
	barrierTwo: semaphore.Semaphore(1),
	numThreads: len(funcs),
	funcs: funcs,
	}
}


func (this *TwoPhaseBarrier) Run() (sem *semaphore.Semaphore) {
	for _, f := range this.funcs {	go f(this) }
	return &this.end // caller waits on this
}