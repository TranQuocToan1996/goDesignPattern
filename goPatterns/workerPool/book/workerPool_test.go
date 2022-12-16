package book

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
	"time"
)

const (
	bufferSize = 100
	timeout    = time.Second * 5
	workers    = 3
	requests   = 10
)

func TestWorkerPool(t *testing.T) {

	var dispatcher IDispatcher = NewDispatcher(bufferSize, timeout)
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	wg.Wait()
	dispatcher.Stop()
}

func Test_Dispatcher(t *testing.T) {

	dispatcher := NewDispatcher(bufferSize, timeout)

	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}

	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				defer wg.Done()
				s, ok := i.(string)
				if !ok {
					t.Fail()
				}
				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) ->
               [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			}}
		dispatcher.MakeRequest(req)
	}

}
