package main

import (
	"fmt"
	"math"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			if i > 0 && j > 0 && s[i-1] == t[j-1] {
				prev := math.Max(float64(dp[i][j-1]), float64(dp[i][j]))
				dp[i][j] = int(math.Max(float64(dp[i-1][j-1]+1), prev))
			} else if i > 0 && j > 0 {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	fmt.Println(dp[n][m])
}
