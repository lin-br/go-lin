package generics_test

import (
	"testing"

	"github.com/lin-br/go-lin/generics"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := generics.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d and want %d, given %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := generics.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d and want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := generics.SumAll([]int{1, 2, 0, 9})
	want := 12

	AssertEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := generics.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		AssertSlicesEqual(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := generics.SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		AssertSlicesEqual(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, generics.Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, generics.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}
