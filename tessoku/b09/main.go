package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([][]int, 1502)
	for i := range x {
		x[i] = make([]int, 1502)
	}

	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Scan(&d)

		x[a][b]++
		x[a][d]--
		x[c][b]--
		x[c][d]++
	}

	for i := 0; i <= 1501; i++ {
		hrz := 0
		for j := 0; j <= 1501; j++ {
			hrz += x[i][j]
			x[i][j] = hrz
		}
	}

	for j := 0; j <= 1501; j++ {
		vrt := 0
		for i := 0; i <= 1501; i++ {
			vrt += x[i][j]
			x[i][j] = vrt
		}
	}

	ans := 0
	for i := 0; i <= 1501; i++ {
		for j := 0; j <= 1501; j++ {
			if x[i][j] > 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)

}
