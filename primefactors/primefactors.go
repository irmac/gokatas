package katas

func primeFactorsOf(n int) []int {
	var factors []int
	if n == 2 {
		factors = append(factors, 2)
	}
	return factors
}
