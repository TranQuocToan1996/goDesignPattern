package gopatterns

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrencySingoton(t *testing.T) {

	const (
		iterationNum = 5000
	)
	numberOfSingleton := 2
	singleton := GetInstance()
	singleton2 := GetInstance()

	for i := 0; i < iterationNum; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())
	var val int

	for val != iterationNum*numberOfSingleton {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()

}
