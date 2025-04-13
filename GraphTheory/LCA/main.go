package main

import "fmt"

const N = 4e4 + 5

var n, m int
var fa [N][30]int // fa[i][k] 表示从 i 向上跳 2 ^ k 格
var depth [N]int  // depth[i] 表示 i 号节点的深度

func LCA(x, y int) int {

}

func main() {
	fmt.Scan(&n, &m)

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		fa[b][0] = a
	}

	// 初始化

	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(LCA(a, b))
	}
}
