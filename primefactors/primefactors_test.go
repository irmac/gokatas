package katas

import "testing"

func TestPrimeFactorsOfOne(t *testing.T) {
	actual := primeFactorsOf(1)
	var expected []int
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfTwo(t *testing.T) {
	actual := primeFactorsOf(2)
	expected := []int{2}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfThree(t *testing.T) {
	actual := primeFactorsOf(3)
	expected := []int{3}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfFour(t *testing.T) {
	actual := primeFactorsOf(4)
	expected := []int{2, 2}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfFive(t *testing.T) {
	actual := primeFactorsOf(5)
	expected := []int{5}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfSix(t *testing.T) {
	actual := primeFactorsOf(6)
	expected := []int{2, 3}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfSeven(t *testing.T) {
	actual := primeFactorsOf(7)
	expected := []int{7}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfEight(t *testing.T) {
	actual := primeFactorsOf(8)
	expected := []int{2, 2, 2}
	assertThat(t, expected, actual)
}

func TestPrimeFactorsOfNine(t *testing.T) {
	actual := primeFactorsOf(9)
	expected := []int{3, 3}
	assertThat(t, expected, actual)
}
