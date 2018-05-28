package problems

import (
	"fmt"
)

type EvenFibonacciNumbers struct{}

func (p *EvenFibonacciNumbers) ID() int {
	return 2
}

func (p *EvenFibonacciNumbers) Text() string {
	return `Each new term in the Fibonacci sequence is generated by adding the
	previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.
`
}

func (p *EvenFibonacciNumbers) Solve() (string, error) {
	limit := uint64(4000000)
	terms := make(chan uint64)
	go p.fibonacci(terms)

	sum := uint64(0)
	for term := range terms {
		if term > limit {
			break
		}
		if term&1 == 0 {
			sum += term
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func (p *EvenFibonacciNumbers) fibonacci(terms chan uint64) {
	vals := []uint64{1, 1}
	for {
		vals[0], vals[1] = vals[1], vals[0]+vals[1]
		terms <- vals[0]
	}
}
