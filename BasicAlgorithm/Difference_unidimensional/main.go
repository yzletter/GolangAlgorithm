package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/799/

func main() {
	in := bufio.NewReader(os.Stdin)
	//out := bufio.NewWriter(os.Stdout)

	var n, q int
	fmt.Scan(&n, &q)

	a := make([]int, n+5)
	sa := make([]int, n+5)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sa[i] += a[i]
		sa[i+1] -= a[i]
	}

	var l, r, c int
	for q > 0 {
		fmt.Fscan(in, &l, &r, &c)
		sa[l] += c
		sa[r+1] -= c
		q--
	}

	for i := 1; i <= n; i++ {
		sa[i] += sa[i-1]
		fmt.Printf("%d ", sa[i])
	}

}
