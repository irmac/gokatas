package katas

func primeFactorsOf(n int) []int {
	var factors []int
	if n > 1 {
		if n%2 == 0 {
			factors = append(factors, 2)
			n /= 2
		}
		if n > 1 {
			factors = append(factors, n)
		}
	}
	return factors
}
