package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/description/800/
func insert(a [][]int, x1, y1, x2, y2, c int) {
	a[x1][y1] += c
	a[x2+1][y1] -= c
	a[x1][y2+1] -= c
	a[x2+1][y2+1] += c
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q int
	fmt.Scan(&n, &m, &q)

	a := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		a[i] = make([]int, m+2)
	}

	var x int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &x)
			insert(a, i, j, i, j, x)
		}
	}

	var x1, y1, x2, y2 int
	for q > 0 {
		fmt.Fscan(in, &x1, &y1, &x2, &y2, &x)

		insert(a, x1, y1, x2, y2, x)
		q--
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] += a[i-1][j] + a[i][j-1] - a[i-1][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}

}
