package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//https://www.acwing.com/problem/content/description/805/

type seg struct {
	l int
	r int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]seg, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].l, &a[i].r)
	}

	sort.Slice(a,
		func(i, j int) bool {
			return a[i].l <= a[j].l
		})

	ans := 0
	st, ed := a[0].l, a[0].r
	for i := 1; i < n; i++ {
		now := a[i]
		if ed < now.l {
			ans++
			st = now.l
			ed = now.r
		} else {
			st = st
			if ed < now.r {
				ed = now.r
			}
		}
	}

	fmt.Println(ans + 1)
}
