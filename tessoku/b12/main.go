package main

import "fmt"

var n float64

func main() {
	fmt.Scan(&n)

	l := 1.0
	r := 100_000_000.0

	for (r > l) && (r-l > 0.001) || (l > r) && (l-r > 0.001) {
		mid := (l + r) / 2.0
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(l)
}

func check(x float64) bool {
	ans := x*x*x + x
	return ans >= n
}
