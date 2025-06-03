package main

import (
	"fmt"
	"math"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = i
		if i == 0 {
			for j := 1; j <= len(t); j++ {
				dp[0][j] = j
			}
			continue
		}
		for j := 1; j <= len(t); j++ {
			rd := 0
			if s[i-1] != t[j-1] {
				rd = 1
			}
			dp[i][j] = int(math.Min(
				float64(dp[i-1][j]+1), // 削除
				math.Min(
					float64(dp[i][j-1]+1),    // 挿入
					float64(dp[i-1][j-1]+rd), // 置換
				),
			))
		}
	}
	fmt.Println(dp[len(s)][len(t)])
}
