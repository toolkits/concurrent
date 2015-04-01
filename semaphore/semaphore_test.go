package concurrent

import (
	"testing"
)

func TestSemaphore(t *testing.T) {
	sema := NewSemaphore(2)
	sema.Acquire()
	sema.Acquire()

	if sema.AvailablePermits() != 0 {
		t.Error("AvailablePermits, AvailablePermits")
	}

	sema.Release()
	if sema.AvailablePermits() != 1 {
		t.Error("AvailablePermits, AvailablePermits")
	}
	sema.Release()
}
