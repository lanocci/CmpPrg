package main

import "fmt"

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
			} else {
				dp[i][j] = false
			}
		}
	}
	if !dp[n][s] {
		fmt.Println(-1)
		return
	}
	p := make([]int, n+1)
	c := 0
	j := s
	for j > 0 {
		i := n // Start from the last item
		for i > 0 {
			if dp[i-1][j] {
				i-- // If we can achieve j without the i-th item, move to the previous item
			} else if j-a[i] >= 0 && dp[i-1][j-a[i]] {
				p[c] = i  // If we can achieve j with the i-th item, record it
				j -= a[i] // Reduce j by the value of the i-th item
				c++
				break // Break to start from the last item again
			} else {
				i-- // If neither condition is met, move to the previous item
			}
		}
	}

	fmt.Println(c)
	for i := c - 1; i >= 0; i-- {
		fmt.Print(p[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
}
