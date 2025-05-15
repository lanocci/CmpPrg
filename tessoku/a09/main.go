package main

import (
	"fmt"
)

func main() {
	var h, w, n int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&n)

	x := make([][]int, 1501)
	for i := range x {
		x[i] = make([]int, 1501)
	}

	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Scan(&d)
		x[a][b]++
		x[a][d+1]--
		x[c+1][b]--
		x[c+1][d+1]++
	}

	for i := 1; i <= h; i++ {
		hrz := 0
		for j := 1; j <= w; j++ {
			hrz += x[i][j]
			x[i][j] = hrz
		}
	}

	for j := 1; j <= w; j++ {
		vrt := 0
		for i := 1; i <= h; i++ {
			vrt += x[i][j]
			x[i][j] = vrt
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if j > 1 {
				fmt.Print(" ")
			}
			fmt.Print(x[i][j])
		}
		fmt.Println("")
	}
}
