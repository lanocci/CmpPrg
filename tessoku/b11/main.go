package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, q int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&x)
		ans := sort.Search(n, func(j int) bool {
			return a[j] >= x
		})
		fmt.Println(ans)
	}
}
