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

	answer := false

	for i := range 1 << n {
		sum := 0
		for j := range n {
			div := 1 << j
			if (i/div)%2 == 1 {
				sum += a[j]
			}
			if sum == k {
				answer = true
			}
		}
	}
	if answer {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
