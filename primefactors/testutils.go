package katas

import "testing"

func assertThat(t *testing.T, expected, actual []int) {
	if !sliceCompare(expected, actual) {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func sliceCompare(a, b []int) bool {

	if bothNil(a, b) {
		return true
	}

	if oneNil(a, b) {
		return false
	}

	if lenNotEqual(a, b) {
		return false
	}

	for i := range a {
		if notEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func bothNil(a, b []int) bool {
	return a == nil && b == nil
}

func oneNil(a, b []int) bool {
	return a == nil || b == nil
}

func lenNotEqual(a, b []int) bool {
	return len(a) != len(b)
}

func notEqual(a, b int) bool {
	return a != b
}
