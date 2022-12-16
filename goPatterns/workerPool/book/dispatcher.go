package book

import "time"

// can launch an injected WorkerLaunchers type in its own LaunchWorker method
type IDispatcher interface {
	LaunchWorker(w WorkerLauncher) // Facade is to hide launching implementation details
	MakeRequest(Request)
	Stop()
}

func NewDispatcher(b int, timeout time.Duration) IDispatcher {
	return &dispatcher{
		inCh:    make(chan Request, b),
		open:    true,
		timeout: timeout,
	}
}

// launch workers in parallel and handle all the possible incoming channels
type dispatcher struct {
	open    bool
	timeout time.Duration
	inCh    chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}
func (d *dispatcher) Stop() {
	if d.open {
		close(d.inCh)
		d.open = false
	}
}
func (d *dispatcher) MakeRequest(r Request) {
	if !d.open {
		d.inCh = make(chan Request)
	}
	select {
	case d.inCh <- r:
	case <-time.After(d.timeout):
		return
	}
}
