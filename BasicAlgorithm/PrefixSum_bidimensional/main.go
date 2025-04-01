package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/description/798/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Scanf("%d %d %d", &n, &m, &q)

	a := make([][]int, n+1)
	sa := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		a[i] = make([]int, m+1)
		sa[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
			sa[i][j] = sa[i-1][j] + sa[i][j-1] - sa[i-1][j-1] + a[i][j]
		}
	}

	var x1, y1, x2, y2 int
	for q > 0 {
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		fmt.Fprintln(out, sa[x2][y2]-sa[x2][y1-1]-sa[x1-1][y2]+sa[x1-1][y1-1])
		q--
	}
}
