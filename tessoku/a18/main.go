package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	dp := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
		dp[i][0] = true
		if i == 0 {
			continue
		}
		for j := 1; j <= s; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			} else if j-a[i] >= 0 && dp[i-1][j-a[i]] {
				dp[i][j] = true
			}
		}
	}
	if dp[n][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
