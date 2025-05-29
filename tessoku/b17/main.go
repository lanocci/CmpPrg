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
	dp[2] = int(math.Abs(float64(a[1] - a[2])))
	for i := 3; i <= n; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+int(math.Abs(float64(a[i]-a[i-1])))), float64(dp[i-2]+int(math.Abs(float64(a[i]-a[i-2]))))))
	}
	p := make([]int, n+1)
	c := 0
	i := n
	for i > 0 {
		if dp[i] == dp[i-1]+int(math.Abs(float64(a[i]-a[i-1]))) || i == 1 {
			p[c] = i
			i--
		} else if i > 1 && dp[i] == dp[i-2]+int(math.Abs(float64(a[i]-a[i-2]))) {
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
