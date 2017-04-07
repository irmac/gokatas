package katas

import "testing"

func TestSumOfPrimeFactorsOfZero(t *testing.T) {
	expected := 0
	actual := factors(0)

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfTwo(t *testing.T) {
	expected := 2
	actual := factors(2)

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfPrimeFactorsOfThree(t *testing.T) {
	expected := 3
	actual := factors(3)

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}

}
