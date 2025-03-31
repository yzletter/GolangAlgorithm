package main

import "fmt"

// https://www.acwing.com/problem/content/789/

func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	tmp := make([]int, r-l+1)
	i, j, k := l, mid+1, 0

	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}

	for j <= r {
		tmp[k] = arr[j]
		k++
		j++
	}

	k = 0
	for i := l; i <= r; i++ {
		arr[i] = tmp[k]
		k++
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n+5)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	MergeSort(arr, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
