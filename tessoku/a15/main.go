package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	uc := 0
	x := make([]int, n)
	copy(x, a)
	sort.Ints(x)
	var tmp int
	for i := 0; i < n; i++ {
		if tmp != x[i] {
			x[uc] = x[i]
			uc++
			tmp = x[i]
		}
	}
	fmt.Printf("a: %v, x: %v", a, x[:uc])

	for i := 0; i < n; i++ {
		y := a[i]
		l := 0
		r := uc - 1
		for l <= r {
			mid := (l + r) / 2
			if x[mid] == y {
				b[i] = mid + 1
				break
			} else if x[mid] < y {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(b[i])
	}

}
