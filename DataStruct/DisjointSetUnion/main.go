package main

import "fmt"

// https://www.acwing.com/problem/content/838/

var p []int

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m)

	p = make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}

	var op string
	for i := 0; i < m; i++ {
		fmt.Scan(&op, &a, &b)

		if op == "M" {
			p[find(a)] = find(b)
		} else {
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}

}
