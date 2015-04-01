package concurrent

import ()

type Semaphore struct {
	bufSize int
	channel chan int8
}

func NewSemaphore(concurrencyNum int) *Semaphore {
	return &Semaphore{channel: make(chan int8, concurrencyNum), bufSize: concurrencyNum}
}

func (this *Semaphore) Acquire() {
	this.channel <- int8(0)
}

// Release前 必须Acquire
func (this *Semaphore) Release() {
	<-this.channel
}

func (this *Semaphore) AvailablePermits() int {
	return this.bufSize - len(this.channel)
}
