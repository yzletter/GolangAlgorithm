package main

// https://leetcode.cn/problems/count-good-numbers/?envType=daily-question&envId=2025-04-13

func countGoodNumbers(n int64) int {
	return qmi(5, (int(n)+1)/2, 1e9+7) * qmi(4, int(n)/2, 1e9+7) % (1e9 + 7)
}

// qmi 返回 a ^ k % p
func qmi(a, k, p int) int {
	res := 1
	for k > 0 {
		if k&1 == 1 {
			res = res * a % p
		}
		k >>= 1
		a = a * a % p
	}
	return res
}
