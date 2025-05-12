package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	m := make([][]int, h+1)
	m[0] = make([]int, w+1)

	for i := 1; i < h; i++ {
		var hrz int
		m[i] = make([]int, w+1)
		for j := 1; j < w; j++ {
			var v int
			fmt.Scan(&v)
			hrz += v
			m[i][j] = hrz
		}
	}

	for i := 2; i < h; i++ {
		for j := 1; j < w; j++ {
			m[i][j] = m[i-1][j] + m[i][j]
		}
	}

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var a, b, c, d int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Scan(&d)

		ans := m[c][d] + m[a-1][b-1] - m[a-1][d] - m[c][b-1]
		fmt.Println(ans)
	}

}
