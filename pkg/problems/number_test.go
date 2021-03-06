package problems

import (
	"testing"
)

func TestRational_ToDecimal(t *testing.T) {
	r := Rational{Numerator: 1, Denominator: 8}
	d := r.Decimal()
	if d.String() != "0.125" {
		t.Errorf("unexpected conversion from 1/8 to %s", d.String())
	}

	r = Rational{Numerator: 1, Denominator: 29}
	d = r.Decimal()
	if d.String() != "0.(0344827586206896551724137931)" {
		t.Errorf("unexpected conversion from 1/29 to %s", d.String())
	}
}

func TestNumDigits(t *testing.T) {
	tests := []struct {
		Value  uint64
		Digits int
	}{
		{13, 2},
		{0, 1},
		{1, 1},
		{9, 1},
		{99999, 5},
		{10000, 5},
	}

	for i, test := range tests {
		if numDigits(test.Value) != test.Digits {
			t.Errorf("test: %d Expected digits for %d: %d got: %d", i, test.Value, test.Digits, numDigits(test.Value))
		}
	}
}

func TestIsPandigital(t *testing.T) {
	tests := []struct {
		Value  uint64
		Result bool
	}{
		{1, true},
		{2, false},
		{746589321, true},
		{465321, true},
		{5465321, false},
	}

	for i, test := range tests {
		if IsPandigital(test.Value) != test.Result {
			t.Errorf("test: %d IsPandigital(%d) Expected result %t: got: %t", i, test.Value, test.Result, IsPandigital(test.Value))
		}
	}
}
