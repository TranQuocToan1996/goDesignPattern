package anhthi

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	fmt.Printf("Even numbers up to 8:\n")
	printEvenNumbers(8)
	fmt.Printf("Even numbers up to 9:\n")
	printEvenNumbers(9)
	fmt.Printf("Error: even numbers up to -1:\n")
	printEvenNumbers(-1)
}
