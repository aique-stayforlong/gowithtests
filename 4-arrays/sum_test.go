package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15
	
		if got != expected {
			t.Errorf("expected %d, got %d from %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	collection1 := []int{1, 2}
	collection2 := []int{0, 9}

	got := SumAll(collection1, collection2)
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	// declaramos el método asignando una función anónima a una variable
	checkSums := func (t testing.TB, got, expected []int) {
		t.Helper()
	
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		collection1 := []int{1, 2}
		collection2 := []int{0, 9}
	
		got := SumAllTails(collection1, collection2)
		expected := []int{2, 9}
	
		checkSums(t, got, expected)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}
