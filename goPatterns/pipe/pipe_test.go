package pipe

import "testing"

const (
	cap = 100
)

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55}}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}

// Same example with fanin fanout example
func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := powerSquare(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh
	return result
}

func generator(max int) <-chan int {
	outChInt := make(chan int, cap)
	go func() {
		defer close(outChInt)
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
	}()
	return outChInt
}

func powerSquare(in <-chan int) <-chan int {
	out := make(chan int, cap)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, cap)
	go func() {
		defer close(out)
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
	}()
	return out
}
