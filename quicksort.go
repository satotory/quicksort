package quicksort

func quickSort(arr []int, first, last int) []int {
	if first < last {
		var p int
		arr, p = parts(arr, first, last)
		arr = quickSort(arr, first, p-1)
		arr = quickSort(arr, p+1, last)
	}
	return arr
}

func parts(arr []int, first, last int) ([]int, int) {
	vlast := arr[last]
	c := first

	for i := c; i < last; i++ {
		if arr[i] < vlast {
			arr[c], arr[i] = arr[i], arr[c]
			c++
		}
	}

	arr[c], arr[last] = arr[last], arr[c]

	return arr, c
}
