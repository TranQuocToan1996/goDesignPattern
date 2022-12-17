package publishsubscriber

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}
func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

func TestWriter(t *testing.T) {
	sub := NewWriterSubscriber(0, os.Stdout)
	msg := "Hello"
	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}
	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	sub.Close()
}

func TestPublisher(t *testing.T) {
	msg := "Hello"

	p := NewPublisher()

	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("Could not assert result"))
			}
			if s != msg {
				t.Fail()
			}
		},
		closeTestingFunc: func() {
			fmt.Println("close testing fn")
		}}
	p.start()
	p.AddSubscriberCh() <- sub
	p.PublishingCh() <- msg
	pubCon, ok := p.(*publisher)
	if len(pubCon.subscribers) != 1 || !ok {
		t.Error("Unexpected number of subscribers")
	}
	p.RemoveSubscriberCh() <- sub
	//Number of subscribers is restored to zero
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}
	p.Stop()
}
