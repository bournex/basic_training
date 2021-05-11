package attach

import (
	"testing"
)

func TestAttachCart(t *testing.T) {
	examples := []struct {
		Input  [][]int
		Expect bool
	}{
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			true,
		},
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
			},
			false,
		},
	}

	for _, v := range examples {
		t.Run("violence solution", func(t *testing.T) {
			res := attackCart(v.Input)
			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}

	for _, v := range examples {
		t.Run("complexity n solution", func(t *testing.T) {
			res := attackCart1(v.Input)
			if res != v.Expect {
				t.Errorf("expect %t, got %t", v.Expect, res)
			}
		})
	}
}

func BenchmarkAttachCart(b *testing.B) {
	examples := []struct {
		Input  [][]int
		Expect bool
	}{
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			true,
		},
	}

	for _, v := range examples {
		b.Run("benchmark violence solution", func(b *testing.B) {
			attackCart(v.Input)
		})
	}

	for _, v := range examples {
		b.Run("benchmark complexity n solution", func(b *testing.B) {
			attackCart1(v.Input)
		})
	}
}

func TestAttachQueen(t *testing.T) {
	examples := []struct {
		Input  [][]int
		Expect bool
	}{
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
			},
			true,
		},
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
			},
			false,
		},
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			true,
		},
	}

	// for i, v := range examples {
	// 	t.Run("violence solution", func(t *testing.T) {
	// 		res := attackQueen(v.Input)
	// 		if res != v.Expect {
	// 			t.Errorf("%d, expect %t, got %t", i, v.Expect, res)
	// 		}
	// 	})
	// }

	for i, v := range examples {
		t.Run("complexity n solution", func(t *testing.T) {
			res := attackQueen1(v.Input)
			if res != v.Expect {
				t.Errorf("%d, expect %t, got %t", i, v.Expect, res)
			}
		})
	}
}

func BenchmarkAttachCart1(b *testing.B) {
	examples := []struct {
		Input  [][]int
		Expect bool
	}{
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			true,
		},
	}

	// for _, v := range examples {
	// 	b.Run("benchmark violence solution", func(b *testing.B) {
	// 		attackQueen(v.Input)
	// 	})
	// }

	for _, v := range examples {
		b.Run("benchmark complexity n solution", func(b *testing.B) {
			attackQueen1(v.Input)
		})
	}
}
