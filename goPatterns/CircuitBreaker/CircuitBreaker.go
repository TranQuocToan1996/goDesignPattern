package circuitbreaker

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"
)

var (
	ErrServiceUnavailable = errors.New("not enough time has passed since the last failure to retry")
)

type State int

const (
	UnknownState State = iota
	FailureState
	SuccessState
)

// a simple counter that records success and failure states of a circuit along with a timestamp and calculates the consecutive number of failures
type Counter interface {
	Count(State)
	ConsecutiveFailures() uint32
	LastActivity() time.Time
	Reset()
}

// https://learn.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10)?redirectedfrom=MSDN
// https://levelup.gitconnected.com/circuit-breaker-example-in-golang-e6459c87eaeb
type CounterImpl struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailure   uint32
}

func (c *CounterImpl) Count(State) {
	// TODO
}

func (c *CounterImpl) ConsecutiveFailures() uint32 {
	// TODO
	return 1
}
func (c *CounterImpl) LastActivity() time.Time {
	// TODO
	return time.Now()
}
func (c *CounterImpl) Reset() {
	// TODO
}

func NewCounter() Counter {
	return &CounterImpl{}
}

// Circuit type is used here to carry deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes
type Circuit func(context.Context) error

func Breaker(c Circuit, failureThreshold uint32) Circuit {
	cnt := NewCounter()

	return func(ctx context.Context) error {
		if cnt.ConsecutiveFailures() >= failureThreshold {
			canRetry := func(cnt Counter) bool {
				backoffLevel := cnt.ConsecutiveFailures() - failureThreshold

				// Calculates when should the circuit breaker resume propagating requests
				// to the service
				shouldRetryAt := cnt.LastActivity().Add(time.Second * 2 << backoffLevel)

				return time.Now().After(shouldRetryAt)
			}

			if !canRetry(cnt) {
				// Fails fast instead of propagating requests to the circuit since
				// not enough time has passed since the last failure to retry
				return ErrServiceUnavailable
			}
		}

		// Unless the failure threshold is exceeded the wrapped service mimics the
		// old behavior and the difference in behavior is seen after consecutive failures
		if err := c(ctx); err != nil {
			cnt.Count(FailureState)
			return err
		}

		cnt.Count(SuccessState)
		return nil
	}
}

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Printf("%s lasted %s", name, elapsed)
}

func BigIntFactorial(x big.Int) *big.Int {
	// Arguments to a defer statement is immediately evaluated and stored.
	// The deferred function receives the pre-evaluated values when its invoked.
	defer Duration(time.Now(), "IntFactorial")

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(&x, one) {
		y.Mul(y, &x)
	}

	return x.Set(y)
}
