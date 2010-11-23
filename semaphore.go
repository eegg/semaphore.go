package semaphore

import "runtime"

type Semaphore uint32

func (this *Semaphore) P() { runtime.Semacquire((*uint32)(this)) }
func (this *Semaphore) V() { runtime.Semrelease((*uint32)(this)) }
func (this *Semaphore) Turnstile() { this.P(); this.V() }

func (this *Semaphore) Do(do func()) { this.P(); do(); this.V() }
