package main

import "fmt"

//https://www.acwing.com/problem/content/description/3/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const N = 1e3 + 5

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	v, w, f := make([]int, N), make([]int, N), make([]int, N)

	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	for i := 1; i <= n; i++ { // 枚举物品
		for j := v[i]; j <= m; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i]) // 是否选当前物品
		}
	}

	fmt.Println(f[m])
}
