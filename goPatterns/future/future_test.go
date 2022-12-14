package future

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		//Timeout!
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(setContext("Hello"))
		wg.Wait()
	})
}
