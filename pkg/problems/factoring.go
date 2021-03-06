package problems

// Given a composite number, determine if factor is a factor.
// If so return true and the quotient of composite/factor
// If not return false and the original composite
func removeFactor(composite uint64, factor uint64) (bool, uint64) {
	if composite < 2 {
		return false, composite
	}
	attempt := composite / factor
	if attempt*factor == composite {
		return true, attempt
	}
	return false, composite
}

func primeFactors(n uint64) map[uint64]int {
	var ok bool
	factor := uint64(2)
	pfs := map[uint64]int{}

	for {
		ok, n = removeFactor(n, factor)
		if ok {
			pfs[factor]++
			if n <= 1 {
				break
			}
		} else {
			factor++
		}
		if n <= 1 {
			break
		}
	}
	return pfs
}

// Find the least common multiple for an array of integers
// Break each integer into its prime factors and
// combine for the lcm
// ex:
//   12, 45 = 2^2*3, 3^2*5 = 2^2*3^2*5
func lcm(ns []uint64) uint64 {
	factors := map[uint64]int{}

	for _, n := range ns {
		pf := primeFactors(n)
		for f, c := range pf {
			if factors[f] < c {
				factors[f] = c
			}
		}
	}

	prod := uint64(1)
	for f, c := range factors {
		for i := 0; i < c; i++ {
			prod *= f
		}
	}
	return prod
}

// Find the greatest common factor for an array of integers
// Break the first integer into prime factors and
// remove if don't match the prime factors of further numbers.
// ex:
//   36, 27 = 2^2*3*2, 3^3 = 3^2 = 9
func gcf(ns []uint64) uint64 {
	factors := primeFactors(ns[0])
	for _, n := range ns[1:] {
		check := primeFactors(n)
		for k, v := range check {
			if factors[k] > v {
				factors[k] = v
			}
		}
		for k, _ := range factors {
			if check[k] == 0 {
				delete(factors, k)
			}
		}
	}
	prod := uint64(1)
	for f, c := range factors {
		for i := 0; i < c; i++ {
			prod *= f
		}
	}
	return prod
}
