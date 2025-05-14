package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)

	//p := make([][]int, n)
	tds := make([][]int, 1501)
	for i := 0; i < 1501; i++ {
		tds[i] = make([]int, 1501)
	}

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		//p[i] = []int{x, y}
		tds[x][y]++
	}
	for i := 1; i <= 1500; i++ {
		hrz := 0
		for j := 1; j <= 1500; j++ {
			hrz += tds[i][j]
			tds[i][j] = hrz
		}
	}

	for i := 2; i <= 1500; i++ {
		for j := 1; j <= 1500; j++ {
			tds[i][j] += tds[i-1][j]
		}
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var a, b, c, d int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Scan(&d)

		ans := tds[c][d] + tds[a-1][b-1] - tds[a-1][d] - tds[c][b-1]
		fmt.Println(ans)
	}
}
