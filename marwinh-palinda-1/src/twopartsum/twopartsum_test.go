package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func TestSumConcurrentCorrectlySumsEvenArray1(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 54

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func TestSumConcurrentCorrectlySumsEvenArray2(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 54 / 2

	actual := ConcurrentSum(arr) / 2

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}
