package katas

func primeFactorsOf(n int) []int {
	var factors []int
	if n > 1 {
		factors = append(factors, 2)
	}
	return factors
}
