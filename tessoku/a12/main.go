package main

import "fmt"

var n, k int
var a []int

func main() {
	fmt.Scan(&n)
	fmt.Scan(&k)

	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	l := 1
	r := 1_000_000_000
	var mid int

	for l < r {
		mid = (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Println(l)
}

func check(x int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += x / a[i]
	}
	return sum >= k
}
