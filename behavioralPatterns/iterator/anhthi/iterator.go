package anhthi

import (
	"fmt"
	"log"
)

// EvenNumberIterator generates even numbers
type EvenNumberIterator struct {
	max       int
	currValue int
	err       error
}

// NewEvenNumberIterator creates new number iterator
func NewEvenNumberIterator(max int) *EvenNumberIterator {
	var err error
	if max < 0 {
		err = fmt.Errorf("'max' is %d, should be >= 0", max)
	}
	return &EvenNumberIterator{
		max:       max,
		currValue: 0,
		err:       err,
	}
}

// Next advances to next even number. Returns false on end of iteration.
func (i *EvenNumberIterator) Next() bool {
	if i.err != nil {
		return false
	}
	i.currValue += 2
	return i.currValue <= i.max
}

// Value returns current even number
func (i *EvenNumberIterator) Value() int {
	if i.err != nil || i.currValue > i.max {
		panic("Value is not valid after iterator finished")
	}
	return i.currValue
}

// Err returns iteration error.
func (i *EvenNumberIterator) Err() error {
	return i.err
}

func printEvenNumbers(max int) {
	iter := NewEvenNumberIterator(max)
	for iter.Next() {
		fmt.Printf("n: %d\n", iter.Value())
	}
	if iter.Err() != nil {
		// TODO: shouldnt fatal here, return err instead and handle it
		log.Fatalf("error: %s\n", iter.Err())
	}
}

func iterateEvenNumbers(max int, cb func(n int) error) error {
	if max < 0 {
		return fmt.Errorf("'max' is %d, must be >= 0", max)
	}
	for i := 2; i <= max; i += 2 {
		err := cb(i)
		if err != nil {
			return err
		}
	}
	return nil
}
