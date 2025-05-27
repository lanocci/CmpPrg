package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = int(math.Abs(float64(a[2] - a[1])))
	for i := 3; i <= n; i++ {
		dp[i] = int(math.Min(float64(dp[i-1])+math.Abs(float64(a[i]-a[i-1])), float64(dp[i-2])+math.Abs(float64(a[i]-a[i-2]))))
	}

	fmt.Println(dp[n])
}
