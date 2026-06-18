package main

import (
	"time"
)

type semaphore struct {
	ch chan struct{}
}

func (sem *semaphore) Acquire() {
	sem.ch <- struct{}{}
}

func (sem *semaphore) Release() {
	select {
	case <-sem.ch:
	default:
	}
}

func (sem *semaphore) TryAcquire() bool {
	select {
	case sem.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (sem *semaphore) AcquireWithTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case sem.ch <- struct{}{}:
		return true
	case <-timer.C:
		return false
	}
}

func (sem *semaphore) Available() int {
	return cap(sem.ch) - len(sem.ch)
}
