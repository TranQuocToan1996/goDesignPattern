package book

import (
	"fmt"
	"testing"
	"time"
)

func TestCommmandQueue(t *testing.T) {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}

func TestCommandPrint(t *testing.T) {
	var (
		timeCommand  CommandPrint = &TimePassed{time.Now()}
		helloCommand CommandPrint = &HelloMessage{}
	)
	time.Sleep(time.Second)
	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}

func TestCommandWithCOR(t *testing.T) {
	second := new(Logger)
	first := Logger{NextChain: second}
	command := &TimePassed{start: time.Now()}
	first.Next(command)
}
