package katas

import "testing"

func TestSumOfPrimeFactorsOfOne(t *testing.T) {
	actual := primeFactorsOf(0)
	var expected []int
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
