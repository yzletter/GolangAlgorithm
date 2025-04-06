package DFS

// https://leetcode.cn/problems/permutations/solutions/2079585/hui-su-bu-hui-xie-tao-lu-zai-ci-jing-que-6hrh/
func permute(nums []int) (ans [][]int) {
	idx, n, cnt := 0, len(nums), 1
	st, path := make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		cnt = cnt * i
	}

	ans = make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		ans[i] = make([]int, 0)
	}

	var dfs func(int)
	dfs = func(u int) {
		if u == n+1 {
			for i := 1; i <= n; i++ {
				ans[idx] = append(ans[idx], path[i])
			}
			idx++
			return
		}

		for i := 1; i <= n; i++ {
			if st[i] == 0 {
				st[i] = 1
				path[u] = nums[i-1]
				dfs(u + 1)
				st[i] = 0
			}
		}
	}
	dfs(1)
	return ans
}
