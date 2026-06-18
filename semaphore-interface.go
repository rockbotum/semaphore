package main

import (
	"time"
)

type Semaphore interface {
	Acquire()

	Release()

	TryAcquire() bool

	AcquireWithTimeout(timeout time.Duration) bool

	Available() int
}

func NewSemaphore(capacity int) *semaphore {
	if capacity <= 0 {
		panic("limit must be > 0")
	}
	return &semaphore{
		ch: make(chan struct{}, capacity),
	}
}
