package BFS

var n, m int

// https://leetcode.cn/problems/number-of-islands/
// 偏移量
var dx, dy = [4]int{0, 1, 0, -1}, [4]int{-1, 0, 1, 0}

type pair struct {
	x, y int
}

func check(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func floodFill(grid [][]byte, i, j int) {
	q := make([]pair, 0)
	q = append(q, pair{i, j})

	for len(q) > 0 {
		hh := q[0]
		q = q[1:]

		xx, yy := hh.x, hh.y
		grid[xx][yy] = '0'

		for i := 0; i < 4; i++ {
			x, y := xx+dx[i], yy+dy[i]
			if check(x, y) && grid[x][y] != '0' {
				grid[x][y] = '0' // 剪枝
				q = append(q, pair{x, y})
			}
		}
	}

}

func numIslands(grid [][]byte) int {
	var ans int
	m, n = len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				floodFill(grid, i, j)
				ans++
			}
		}
	}
	return ans
}
