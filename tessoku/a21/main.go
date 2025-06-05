package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n+2)
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Scan(&p[i])
		fmt.Scan(&a[i])
	}
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}
	dp[1][n] = 0
	for len := n - 2; len >= 0; len-- {
		for l := 1; l <= n-len; l++ {
			r := l + len
			s1 := 0
			if l <= p[l-1] && p[l-1] <= r {
				s1 = a[l-1]
			}
			s2 := 0
			if p[r+1] <= r && p[r+1] >= l {
				s2 = a[r+1]
			}
			if l == 1 {
				dp[l][r] = dp[l][r+1] + s2
			} else if r == n {
				dp[l][r] = dp[l-1][r] + s1
			} else {
				dp[l][r] = int(math.Max(float64(s1), float64(s2)))
			}

		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dp[i][i] >= ans {
			ans = dp[i][i]
		}
	}
	fmt.Println(ans)
}
