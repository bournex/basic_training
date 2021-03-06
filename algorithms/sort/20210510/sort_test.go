package sort0510

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	getInput := func(n int) []int {
		return rand.Perm(n)
	}
	getExpect := func(n int) []int {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i
		}
		return s
	}

	testCases := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{0, 1},
			[]int{0, 1},
		},
		{
			[]int{4, 7, 5, 1, 0, 8, 9, 2, 3, 6},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			getInput(100),
			getExpect(100),
		},
		{
			getInput(500),
			getExpect(500),
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)

			insert_sort(input)
			for j, v1 := range input {
				if v1 != v.Expect[j] {
					t.Errorf("insert sort input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					break
				}
			}
			copy(input, v.Input)

			select_sort(input)
			for j, v1 := range input {
				if v1 != v.Expect[j] {
					t.Errorf("select sort input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					break
				}
			}
			copy(input, v.Input)

			merge_sort(input)
			for j, v1 := range input {
				if v1 != v.Expect[j] {
					t.Errorf("merge sort input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					break
				}
			}
			copy(input, v.Input)

			quick_sort(input)
			for j, v1 := range input {
				if v1 != v.Expect[j] {
					t.Errorf("quick sort input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					break
				}
			}
			copy(input, v.Input)
		})
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	Input := rand.Perm(10000)
	input := make([]int, 10000)

	b.Run("insert sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, Input)
			insert_sort(input)
		}
	})
	b.Run("select sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, Input)
			select_sort(input)
		}
	})
	b.Run("merge sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, Input)
			merge_sort(input)
		}
	})
	b.Run("quick sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, Input)
			quick_sort(input)
		}
	})
	b.Run("quick sort old", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(input, Input)
			quick_sort_old(input)
		}
	})
}
