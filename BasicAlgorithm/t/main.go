package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	l, r := 0.0, float64(x)

	for r-l > 1e-8 {
		mid := (l + r) / 2

		if mid*mid > float64(x) && mid*mid-float64(x) < 1e-4 {
			fmt.Println(mid)
			return
		}

		if mid*mid > float64(x) && mid*mid-float64(x) < 1e-4 {
			r = mid
		} else {
			l = mid
		}
		fmt.Println(mid)
	}

}
