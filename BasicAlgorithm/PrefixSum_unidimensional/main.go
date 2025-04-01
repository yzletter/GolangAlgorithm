package main

import "fmt"

// https://www.acwing.com/problem/content/797/
func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	a := make([]int, n+5)
	sa := make([]int, n+5)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		sa[i] = sa[i-1] + a[i]
	}

	var l, r int
	for m > 0 {
		fmt.Scanf("%d %d", &l, &r)
		fmt.Println(sa[r] - sa[l-1])
		m--
	}
}
