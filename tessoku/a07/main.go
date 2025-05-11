package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d)
	fmt.Scan(&n)

	b := make([]int, d)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l)
		fmt.Scan(&r)

		b[l-1]++
		b[r]--
	}

	tmp := 0
	for _, v := range b {
		tmp += v
		fmt.Println(tmp)
	}
}
