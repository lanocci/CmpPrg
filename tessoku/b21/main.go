package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 同じインデックスであれば、長さ1の回文となる
	}
	for len := 2; len <= n; len++ { // 長さ1は調査ずみなので、長さ2から短い順に調べる
		for i := 0; i <= n-len; i++ { // 左端のインデックスをiとする
			j := i + len - 1  // 右端のインデックスをjとする
			if s[i] == s[j] { // 左右のインデックスで文字が同じだったら、左端、右端の2文字をくっつけて+2文字の回文が作れる
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1]))) // 左端+1もしくは右端-1の回文のうち、長い方を採用
			}
		}
	}
	fmt.Println(dp[0][n-1])
}
