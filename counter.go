package counter

import "./semaphore"

type Counter struct {
	Mutex semaphore.Semaphore
	Value int
}

func (this *Counter) Do(f func()) { this.Mutex.Do(f) }

func (this *Counter) Add(amt int) {
	this.Do( func() { this.Value += amt } )
}

func (this *Counter) AddAnd(amt int, eq int, f func()) {
	this.Do(func() {
		this.Value += amt
		if this.Value == eq { f() }
	})
}

func (this *Counter) Inc() { this.Add(1) }
func (this *Counter) Dec() { this.Add(-1) }

func (this *Counter) IncAnd(eq int, f func()) { this.AddAnd(1, eq, f) }
func (this *Counter) DecAnd(eq int, f func()) { this.AddAnd(-1, eq, f) }

func NewCounter() Counter { return Counter {Mutex: 1} }