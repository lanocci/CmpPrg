package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	r := make([]int, n+1)

	for i := 1; i < n; i++ {
		if i == 1 {
			r[i] = 1
		} else {
			r[i] = r[i-1]
		}
		for r[i] < n && a[r[i]+1]-a[i] <= k {
			r[i]++
		}
	}

	answer := 0
	for i := 1; i < n; i++ {
		answer += r[i] - i
	}
	fmt.Println(answer)
}
