package main

import "fmt"

// https://www.acwing.com/problem/content/792/
func main() {
	var n float64
	fmt.Scanln(&n)

	l, r := -1000.00, 1000.00
	for r-l > 1e-7 {
		mid := (l + r) / 2
		if mid*mid*mid > n {
			r = mid
		} else {
			l = mid
		}
	}

	fmt.Printf("%.6f", l)
}
