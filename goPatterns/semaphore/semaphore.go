package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

// Semaphore contains the behavior of a semaphore that can be acquired and/or released.
// semaphore: a construct which can be used to constrain or control access to a shared resource access via multiple threads or goroutines
type Semaphore interface {
	Acquire() error
	Release() error
}

// A semaphore is a synchronization pattern/primitive that imposes mutual exclusion on a limited number of resources
type Implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *Implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *Implementation) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Semaphore {
	return &Implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
