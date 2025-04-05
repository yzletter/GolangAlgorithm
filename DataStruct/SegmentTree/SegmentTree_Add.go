package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var a, tr, lazy []int

// 建树
func buildTree(st int, ed int, u int) {
	if st == ed {
		tr[u] = a[st]
		return
	}

	mid := (st + ed) / 2
	buildTree(st, mid, u<<1)
	buildTree(mid+1, ed, u<<1|1)

	pushUp(u)
}

func update(st, ed, u, x int) {
	lazy[u] += x
	tr[u] += (ed - st + 1) * x
}

// 查询区间和
func queryAns(st, ed, u, l, r int) int {
	if st >= l && ed <= r {
		return tr[u]
	}

	pushDown(st, ed, u)
	mid := (st + ed) >> 1
	ans := 0

	if mid >= l {
		ans += queryAns(st, mid, u<<1, l, r)
	}
	if mid+1 <= r {
		ans += queryAns(mid+1, ed, u<<1|1, l, r)
	}
	return ans
}

// 区间修改
func modifyTree(st, ed, u, l, r, x int) {
	if st >= l && ed <= r {
		update(st, ed, u, x)
		return
	}

	pushDown(st, ed, u)

	mid := (st + ed) >> 1
	if mid >= l {
		modifyTree(st, mid, u<<1, l, r, x)
	}
	if mid+1 <= r {
		modifyTree(mid+1, ed, u<<1|1, l, r, x)
	}

	pushUp(u)
}

func pushUp(u int) {
	tr[u] = tr[u<<1] + tr[u<<1|1]
}

func pushDown(st, ed, u int) {
	if lazy[u] == 0 {
		return
	}

	mid := (st + ed) >> 1
	update(st, mid, u<<1, lazy[u])
	update(mid+1, ed, u<<1|1, lazy[u])
	pushUp(u)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanf(in, "%d %d", &n, &m)

	a = make([]int, n+5)
	tr = make([]int, n*4)
	lazy = make([]int, n*4)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	buildTree(1, n, 1)

	var op, l, r, x int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &op)
		if op == 1 {
			fmt.Fscan(in, &l, &r, &x)
			modifyTree(1, n, 1, l, l, x)
		} else {
			fmt.Fscan(in, &l, &r)
			fmt.Println(queryAns(1, n, 1, l, r))
		}
	}
}
