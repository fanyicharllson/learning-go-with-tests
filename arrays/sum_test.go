package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of fixed arrays", func(t *testing.T) {
		num := []int{1, 2, 3, 4, 5}

		got := Sum(num)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, num)
		}
	})
	t.Run("collection of  any size", func(t *testing.T) {
		num := []int{1, 2, 3}

		got := Sum(num)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, num)
		}
	})
	t.Run("computing multi slice", func(t *testing.T) {
		got := SumMulti([]int{1, 2}, []int{0, 9}, []int{9, 3})
		want := []int{3, 9, 12}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("slice with total tails", func(t *testing.T) {
		got := SumTotalTails([]int{1, 2}, []int{0, 9}, []int{9, 3})
		want := []int{2, 9, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
