package of10

import "testing"

func TestFib(t *testing.T) {
	examples := []struct {
		Input  int
		Expect int
	}{
		{
			95,
			407059028,
		},
		{
			40,
			102334155,
		},
		{
			100,
			687995182,
		},
		{
			0,
			0,
		},
		{
			3,
			2,
		},
	}

	for _, v := range examples {
		res := fib(v.Input)
		if res != v.Expect {
			t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
		}
	}

}

func BenchmarkFib(b *testing.B) {
	benchmark := []struct {
		Input int
	}{
		{
			1000,
		},
		{
			100000,
		},
		{
			100000000,
		},
	}
	for _, v := range benchmark {
		b.Run("benchmark", func(b *testing.B) {
			fib(v.Input)
		})
	}
}
