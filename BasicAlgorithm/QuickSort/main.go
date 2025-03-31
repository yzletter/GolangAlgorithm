package main

import "fmt"

// https://www.acwing.com/problem/content/787/
func QuickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	x, i, j := arr[(l+r)/2], l-1, r+1

	for i < j {
		for {
			i++
			if arr[i] >= x {
				break
			}
		}

		for {
			j--
			if arr[j] <= x {
				break
			}
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n+5)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	QuickSort(arr, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
