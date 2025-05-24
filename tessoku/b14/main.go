package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	m := n / 2
	o := n - m

	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, o)
	for i := 0; i < o; i++ {
		fmt.Scan(&b[i])
	}

	p := make([]int, 1<<m)
	q := make([]int, 1<<o)

	for i := 1; i < len(p); i++ {
		for j := 0; j < m; j++ {
			// i をビットで表現し、j番目のビットが立っているかを確認
			// &1はビット演算で、iのj番目のビットが1かどうかをチェック
			if (i>>j)&1 == 1 {
				p[i] += a[j]
			}
		}
		if p[i] == k {
			fmt.Println("Yes")
			return
		}
	}
	sort.Ints(p)

	for i := 1; i < len(q); i++ {
		for j := 0; j < o; j++ {
			if (i>>j)&1 == 1 {
				q[i] += b[j]
			}
		}
		if q[i] == k {
			fmt.Println("Yes")
			return
		}
	}
	sort.Ints(q)

	for i := 1; i < len(p); i++ {
		x := k - p[i]
		l := 0
		r := len(q) - 1
		for l <= r {
			mid := (l + r) / 2
			if q[mid] == x {
				fmt.Println("Yes")
				return
			} else if q[mid] < x {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	fmt.Println("No")

}
