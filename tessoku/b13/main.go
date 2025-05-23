package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	r := make([]int, n)

	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n; j++ {
			if i == j {
				tmp = a[j]
			} else {
				tmp += a[j]
			}
			if tmp > k {
				break
			}
			r[i]++
		}
	}

	answer := 0
	for i := 0; i < n; i++ {
		answer += r[i]
	}
	fmt.Println(answer)
}
