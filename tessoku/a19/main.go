package main

import (
	"fmt"
	"math"
)

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	a := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = make([]int, 2)
		fmt.Scan(&a[i][0], &a[i][1])
	}
	dp := make([][]int, n+1)
	max := 0
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
		dp[i][0] = 0
		for j := 1; j <= w; j++ {
			if i == 0 {
				dp[i][j] = -1000
				break
			}
			if j < a[i][0] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-a[i][0]]+a[i][1])))
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	fmt.Println(max)
}
