package problems

import (
	"fmt"
	"math"
)

type TriangularPentagonalAndHexagonal struct{}

func (p *TriangularPentagonalAndHexagonal) ID() int {
	return 45
}

func (p *TriangularPentagonalAndHexagonal) Text() string {
	return `Triangle, pentagonal, and hexagonal numbers are
generated by the following formulae:

Triangle    T(n)=n(n+1)/2    1, 3, 6, 10, 15, ...
Pentagonal  P(n)=n(3n−1)/2   1, 5, 12, 22, 35, ...
Hexagonal   H(n)=n(2n−1)     1, 6, 15, 28, 45, ...

It can be verified that T(285) = P(165) = H(143) = 40755.

Find the next triangle number that is also pentagonal and hexagonal.
`
}

func (p *TriangularPentagonalAndHexagonal) Solve() (string, error) {

	H := func(n uint64) uint64 {
		return n * (2*n - 1)
	}
	P := func(n uint64) uint64 {
		return n * (3*n - 1) / 2
	}
	T := func(n uint64) uint64 {
		return n * (n + 1) / 2
	}

	quatraticFormula := func(a, b, c float64) []float64 {
		// discriminant
		dis := b*b - 4*a*c
		return []float64{
			(-b + math.Sqrt(dis)) / 2 / a,
			(-b - math.Sqrt(dis)) / 2 / a,
		}
	}

	isPentagon := func(p uint64) bool {
		// 0 = 3*n^2 - n - 2p
		a := float64(3)
		b := float64(-1)
		c := -2 * float64(p)
		ns := quatraticFormula(a, b, c)
		for _, n := range ns {
			if n >= 0 {
				if P(uint64(n)) == p {
					return true
				}
			}
		}
		return false
	}

	isTriangle := func(t uint64) bool {
		// 0 = n^2 + n - 2t
		a := float64(1)
		b := float64(1)
		c := -2 * float64(t)
		ns := quatraticFormula(a, b, c)
		for _, n := range ns {
			if n >= 0 {
				if T(uint64(n)) == t {
					return true
				}
			}
		}
		return false
	}

	done := make(chan bool)
	hex := make(chan uint64, 100)
	pent := make(chan uint64, 100)

	go func() {
		defer close(hex)
		n := uint64(144)
		for {
			select {
			case <-done:
				return
			default:
			}
			hex <- H(n)
			n++
		}
	}()

	go func() {
		for h := range hex {
			if isPentagon(h) {
				pent <- h
			}
		}
		close(pent)
	}()

	for p := range pent {
		if isTriangle(p) {
			return fmt.Sprintf("%d", p), nil
		}
	}
	close(done)

	return fmt.Sprintf("%d", 0), nil
}
