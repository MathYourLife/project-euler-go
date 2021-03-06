package problems

import (
	"fmt"
)

type BigInt struct {
	n []int
}

func NewBigInt(n int) *BigInt {
	b := &BigInt{
		n: []int{n},
	}
	b.Regroup()
	return b
}

func (b *BigInt) Equals(n *BigInt) bool {
	b.Regroup()
	n.Regroup()
	if len(b.n) != len(n.n) {
		return false
	}
	for i := 0; i < len(b.n); i++ {
		if b.n[i] != n.n[i] {
			return false
		}
	}
	return true
}

func (b *BigInt) LessThan(n *BigInt) bool {
	b.Regroup()
	n.Regroup()
	if len(b.n) < len(n.n) {
		return true
	} else if len(b.n) > len(n.n) {
		return false
	}
	for i := len(b.n) - 1; i >= 0; i-- {
		if b.n[i] < n.n[i] {
			return true
		} else if b.n[i] > n.n[i] {
			return false
		}
	}
	return false
}

func (b *BigInt) Regroup() {
	for {
		regrouped := false
		for i := len(b.n) - 1; i >= 0; i-- {
			for {
				if b.n[i] >= 0 {
					break
				}
				b.n[i+1]--
				b.n[i] += 10
				regrouped = true
			}
		}
		if !regrouped {
			break
		}
	}
	// Remove leading zeros
	for i := len(b.n) - 1; i >= 0; i-- {
		if b.n[i] > 0 {
			b.n = b.n[0 : i+1]
			break
		}
	}
	for i := 0; i < len(b.n); i++ {
		if b.n[i] > 9 {
			if i == len(b.n)-1 {
				b.n = append(b.n, 0)
			}
			regroup := b.n[i] / 10
			b.n[i+1] += regroup
			b.n[i] = b.n[i] % 10
		}
	}
}

func (b *BigInt) Mul(f int) {
	for i, _ := range b.n {
		b.n[i] *= f
	}
}

func (b *BigInt) MulBigInt(n *BigInt) *BigInt {
	prod := make([]int, len(b.n)+len(n.n)-1)
	for i := 0; i < len(b.n); i++ {
		for j := 0; j < len(n.n); j++ {
			prod[i+j] += b.n[i] * n.n[j]
		}
	}
	return &BigInt{
		n: prod,
	}
}

func (b *BigInt) AddBigInt(n *BigInt) {
	for i := 0; i < len(n.n) && i < len(b.n); i++ {
		b.n[i] += n.n[i]
	}
	for i := len(b.n); i < len(n.n); i++ {
		b.n = append(b.n, n.n[i])
	}
}

func (b *BigInt) SubBigInt(n *BigInt) {
	for i := 0; i < len(n.n) && i < len(b.n); i++ {
		b.n[i] -= n.n[i]
	}
}

func (b *BigInt) Copy() *BigInt {
	n := make([]int, len(b.n))
	copy(n, b.n)
	bigCopy := &BigInt{
		n: n,
	}
	return bigCopy
}

func (b *BigInt) String() string {
	b.Regroup()
	s := ""
	for _, d := range b.n {
		s = fmt.Sprintf("%d%s", d, s)
	}
	return s
}

func divisors(n uint64) []uint64 {
	pf := primeFactors(n)

	divisors := []uint64{1}
	for factor, exponent := range pf {
		multiple := uint64(1)
		newDivisors := []uint64{}
		for e := 1; e <= exponent; e++ {
			multiple *= factor
			for _, divisor := range divisors {
				newDivisors = append(newDivisors, divisor*multiple)
			}
		}
		divisors = append(divisors, newDivisors...)
	}
	return divisors
}

func properDivisors(n uint64) []uint64 {
	d := divisors(n)
	return d[:len(d)-1]
}
