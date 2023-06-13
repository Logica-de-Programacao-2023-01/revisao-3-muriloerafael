package q3

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	waysClimb := make([]int, n+1)
	waysClimb[1] = 1
	waysClimb[2] = 2

	for i := 3; i <= n; i++ {
		waysClimb[i] = waysClimb[i-1] + waysClimb[i-2]
	}

	return waysClimb[n]
}
