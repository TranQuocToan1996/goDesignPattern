package semaphore

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSemaphoreWithTimeouts(t *testing.T) {
	defer runTime(time.Now())
	tickets, timeout := 1, 3*time.Second
	s := New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		t.Fatal(err)
	}

	// Do important work

	if err := s.Release(); err != nil {
		t.Fatal(err)
	}
}
func TestSemaphore(t *testing.T) {
	defer runTime(time.Now())
	tickets := 0
	s := New(tickets, 0)

	if err := s.Acquire(); err != nil {
		if err != ErrNoTickets {
			t.Fatal(err)
		}
		log.Println("No tickets left, can't work :(", err)
	}
}

func TestSemaphore_LongRequest(t *testing.T) {
	defer runTime(time.Now())
	totProcess := 100
	sem := New(10, 3)
	doneC := make(chan bool, 1)
	for i := 0; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totProcess {
				doneC <- true
			}
		}(i)
	}
	<-doneC
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second)
}

func runTime(begin time.Time) {
	log.Printf("%v", time.Since(begin))
}
