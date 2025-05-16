package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n)
	a := make([]int, n+1)
	p := make([]int, n+1)
	q := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		if a[i] > p[i-1] {
			p[i] = a[i]
		} else {
			p[i] = p[i-1]
		}
	}
	q[n] = a[n]
	for i := n - 1; i >= 1; i-- {
		if a[i] > q[i+1] {
			q[i] = a[i]
		} else {
			q[i] = q[i+1]
		}
	}

	fmt.Scan(&d)

	for i := 0; i < d; i++ {
		var l, r int
		fmt.Scan(&l)
		fmt.Scan(&r)
		if p[l-1] > q[r+1] {
			fmt.Println(p[l-1])
		} else {
			fmt.Println(q[r+1])
		}
	}
}
