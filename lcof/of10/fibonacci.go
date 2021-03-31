package of10

// input:	0,1,2,3,4,5,6,...
// output:	0,1,1,2,3,5,8,...

func fib(n int) int {
	if n < 2 {
		return n
	}

	first, second := 0, 1

	for i := 2; i <= n; i++ {
		first, second = second%1000000007, (first+second)%1000000007
	}

	return second % 1000000007
}
