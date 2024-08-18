package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test fixed number of time", func(t *testing.T) {
		got := Sum([]int{5, 5, 5, 5, 5})
		expected := 25

		if got != expected {
			t.Errorf("expected '%d', got '%d'", expected, got)
		}
	})
	t.Run("test variable number of time", func(t *testing.T) {
		got := Sum([]int{5, 5, 5})
		expected := 15

		if got != expected {
			t.Errorf("expected '%d', got '%d'", expected, got)
		}
	})

}

func TestSumAll(t *testing.T) {
	execpted := []int{3, 5}
	sum := SumAll([]int{1, 2}, []int{2, 3})

	if !reflect.DeepEqual(sum, execpted) {
		t.Errorf("Expected: %v, Actual: %v", execpted, sum)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		expected := []int{2, 8}
		sum := SumAllTails([]int{2, 2}, []int{0, 8})
		checkSums(t, sum, expected)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		expected := []int{0, 8}
		sum := SumAllTails([]int{}, []int{0, 8})
		checkSums(t, sum, expected)
	})

}
