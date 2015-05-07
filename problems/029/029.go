/*
Distinct powers

Consider all integer combinations of a^b for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

    2^2=4,  2^3=8,   2^4=16,  2^5=32
    3^2=9,  3^3=27,  3^4=81,  3^5=243
    4^2=16, 4^3=64,  4^4=256, 4^5=1024
    5^2=25, 5^3=125, 5^4=625, 5^5=3125

If they are then placed in numerical order, with any repeats removed,
we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by
a^b for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/primefactorization"
	"sort"
)

type Exponent struct {
	base     uint64
	exponent uint64
	pf       *primefactorization.PrimeFactorization
}

func NewExponent(base int, exponent int) *Exponent {
	return &Exponent{
		base:     uint64(base),
		exponent: uint64(exponent),
		pf:       primefactorization.NewPrimeFactorization(),
	}
}

func (e *Exponent) Simplify() string {
	s := ""
	m := e.pf.Of(e.base)

	keys := []uint64{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Sort(intArray(keys))

	for _, k := range keys {
		s += fmt.Sprintf("%d^%d*", k, e.exponent*m[k])
	}
	return s
}

type Interface interface {
	Len() uint64
	Less(i, j uint64) bool
	Swap(i, j uint64)
}
type intArray []uint64

func (s intArray) Len() int           { return len(s) }
func (s intArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

func main() {

	N := 100
	d := map[string]bool{}

	for a := 2; a <= N; a++ {
		for b := 2; b <= N; b++ {
			n := NewExponent(a, b)
			d[n.Simplify()] = true
			n = NewExponent(b, a)
			d[n.Simplify()] = true
		}
	}
	fmt.Println(len(d))
}
