package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)

	a := make([]int, n)
	tmp := 0
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > 0 {
			tmp++
		}
		sum[i+1] = tmp
	}
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l)
		fmt.Scan(&r)

		wc := sum[r] - sum[l-1]
		lc := r - l + 1 - wc
		if wc > lc {
			fmt.Println("win")
		} else if wc == lc {
			fmt.Println("draw")
		} else {
			fmt.Println("lose")
		}
	}
}
