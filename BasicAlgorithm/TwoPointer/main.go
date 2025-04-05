package TwoPointer

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mp := make(map[byte]int, n)

	ans := 0
	for i, j := 0, 0; j < n; j++ {
		mp[s[j]]++
		for mp[s[j]] > 1 {
			mp[s[i]]--
			i++
		}

		ans = max(ans, j-i+1)
	}

	return ans
}
