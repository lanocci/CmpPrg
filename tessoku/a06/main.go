// prefix sum

package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	fmt.Scan(&q)

	a := make([]int, n)
	s := make([]int, n+1)
	tmp := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		tmp += a[i]
		s[i+1] = tmp
	}

	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scan(&l)
		fmt.Scan(&r)
		fmt.Println(s[r] - s[l-1])
	}
}
