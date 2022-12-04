package book

import (
	"fmt"
	"time"
)

type CommandPrint interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

// Combine with COR
type ChainLogger interface {
	Next(CommandPrint)
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c CommandPrint) {
	time.Sleep(time.Second)
	fmt.Printf("Elapsed time from creation: %s\n", c.Info())
	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}
