package katas

import "testing"

func TestSumOfPrimeFactorsOfOne(t *testing.T) {
	actual := primeFactorsOf(1)
	var expected []int
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfTwo(t *testing.T) {
	actual := primeFactorsOf(2)
	expected := []int{2}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfThree(t *testing.T) {
	actual := primeFactorsOf(3)
	expected := []int{3}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfFour(t *testing.T) {
	actual := primeFactorsOf(4)
	expected := []int{2, 2}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
