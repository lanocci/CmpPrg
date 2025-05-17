package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	found := false
	l := 0
	r := n
	p := (l + r) / 2

	for !found {
		if a[p] > x {
			r = p - 1
			p = (l + r) / 2
		} else if a[p] < x {
			l = p + 1
			p = (l + r) / 2
		} else {
			found = true
		}
	}
	fmt.Println(p)
}
