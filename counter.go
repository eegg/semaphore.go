package counter

import "./semaphore"

type Counter struct {
	Mutex semaphore.Semaphore
	Value uint32
}

func (this *Counter) Do(f func()) { this.Mutex.Do(f) }

func (this *Counter) Inc() {
	this.Do( func() { this.Value++ })
}

func (this *Counter) Dec() {
	this.Do( func() { this.Value-- })
}

func (this *Counter) IncAnd(eq uint32, f func()) {
	this.Do(func() {
		this.Value++
		if this.Value == eq { f() }
	})
}

func (this *Counter) DecAnd(eq uint32, f func()) {
	this.Do(func() {
		this.Value--
		if this.Value == eq { f() }
	})
}