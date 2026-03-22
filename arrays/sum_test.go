package arrays

import "testing"
import "slices"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("got '%d' but expected '%d', given %v", got, expected, numbers)
		}
	})
	t.Run("collection of 3 number", func(t *testing.T) {
		numbers := []int{5, 6, 7}
		got := Sum(numbers)
		expected := 18
		if got != expected {
			t.Errorf("got '%d' but expected '%d', given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v but expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but expected %v", got, want)
		}
	}
	t.Run("tail sum with non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("tail sum with non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 2, 5, 9})
		want := []int{0, 16}
		checkSums(t, got, want)
	})
}
