package main

import (
	"fmt"
	"sort"
)

var p, q []int

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}

	p = make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p[i*n+j] = a[i] + b[j]
		}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	q = make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			q[i*n+j] = c[i] + d[j]
		}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i] < q[j]
	})

	for i := 0; i < n*n; i++ {
		x := k - p[i]
		l := 0
		r := len(q) - 1
		for l <= r {
			m := (l + r) / 2
			if q[m] == x {
				fmt.Println("Yes")
				return
			} else if q[m] < x {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	fmt.Println("No")

}
