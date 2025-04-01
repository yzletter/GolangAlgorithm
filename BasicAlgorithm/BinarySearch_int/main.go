package main

import "fmt"

// https://www.acwing.com/problem/content/791/
func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	arr := make([]int, n+5)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var x int
	for q > 0 {
		fmt.Scanln(&x)

		l, r := 0, n-1

		for l < r {
			mid := (l + r) / 2
			if arr[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		q--
		if arr[l] != x {
			fmt.Println("-1 -1")
			continue
		}

		fmt.Print(l)

		l, r = 0, n-1
		for l < r {
			mid := (l + r + 1) / 2
			if arr[mid] > x {
				r = mid - 1
			} else {
				l = mid
			}
		}

		fmt.Printf(" %d\n", l)

	}
}
