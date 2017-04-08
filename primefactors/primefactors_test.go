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

func TestSumOfPrimeFactorsOfFive(t *testing.T) {
	actual := primeFactorsOf(5)
	expected := []int{5}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfSix(t *testing.T) {
	actual := primeFactorsOf(6)
	expected := []int{2, 3}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfSeven(t *testing.T) {
	actual := primeFactorsOf(7)
	expected := []int{7}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfEight(t *testing.T) {
	actual := primeFactorsOf(8)
	expected := []int{2, 2, 2}
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfNine(t *testing.T) {

	actual := primeFactorsOf(9)
	expected := []int{3, 3}

	if !sliceCompare(expected, actual) {

		t.Fatalf("Expected %d but got %d", expected, actual)

	}
}
