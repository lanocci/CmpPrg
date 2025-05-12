package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	fmt.Scan(&n)

	b := make([]int, t+1)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l)
		fmt.Scan(&r)

		b[l]++
		b[r]--
	}

	tmp := 0
	for i := 0; i < t; i++ {
		tmp += b[i]
		fmt.Println(tmp)
	}
}
