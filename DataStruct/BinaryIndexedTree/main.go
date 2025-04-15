package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/1266/
const N = 1e5

var n, m int
var tr = make([]int, N+1)

func lowBit(x int) int {
	return x & -x
}

func update(x, a int) {
	for i := x; i <= n; i += lowBit(i) {
		tr[i] += a
	}
}

func query(x int) int {
	res := 0
	for i := x; i >= 1; i -= lowBit(i) {
		res += tr[i]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanf(in, "%d %d", &n, &m)

	var x int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x)
		update(i, x)
	}

	var k, a, b int
	for m > 0 {
		fmt.Fscan(in, &k, &a, &b)
		if k == 1 {
			update(a, b)
		} else {
			fmt.Println(query(b) - query(a-1))
		}
		m--
	}
}
