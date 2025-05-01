package main

// GCD 返回 a, b 的最大公因数
func GCD(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// LCM 返回 a, b 的最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
