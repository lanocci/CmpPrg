package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, n+1)
	for i := 3; i <= n; i++ {
		fmt.Scan(&b[i])
	}
	dp := make([]int, n+1)
	dp[2] = a[2]
	for i := 3; i <= n; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+a[i]), float64(dp[i-2]+b[i])))
	}

	p := make([]int, n)
	c := 0
	i := n
	for i > 0 {
		if dp[i] == dp[i-1]+a[i] {
			p[c] = i
			i--
		} else if dp[i] == dp[i-2]+b[i] {
			p[c] = i
			i -= 2
		}
		c++
	}

	fmt.Println(c)
	for j := c - 1; j >= 0; j-- {
		fmt.Print(p[j])
		if j > 0 {
			fmt.Print(" ")
		}
	}
}
