package gopatterns

import (
	"log"
	"time"
)

var (
	addCh      = make(chan bool)
	getCountCh = make(chan chan int)
	quitCh     = make(chan bool)
)

const (
	timeout = time.Second * 300
)

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}
func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			ticket := time.After(timeout)
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			case <-ticket:
				log.Fatalf("end program due to exceeds timeout %v", timeout)
			}
		}
	}(addCh, getCountCh, quitCh)
}
