package faninfanout

import (
	"fmt"
	"runtime"
	"testing"
)

func TestFaninFanout(t *testing.T) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers = append(numbers, i)
	}

	pipe := generatePipeLine(numbers)

	numCPU := runtime.NumCPU()

	var inputsFanin = []<-chan int{}

	for i := 0; i < numCPU; i++ {
		inputsFanin = append(inputsFanin, fanout(pipe, square))
	}
	var sum int
	for val := range fanin(inputsFanin...) {
		sum += val
	}
	fmt.Println(sum)

}

func square(num int) int {
	return num * num
}

func generatePipeLine(s []int) <-chan int {
	pipe := make(chan int)

	go func() {
		defer close(pipe)
		for _, num := range s {
			pipe <- num
		}
	}()

	return pipe
}

func fanout(in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- f(num)
		}
	}()
	return out
}

func fanin(inputs ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		defer close(in)
		for _, ch := range inputs {
			for square := range ch {
				in <- square
			}
		}
	}()
	return in
}
