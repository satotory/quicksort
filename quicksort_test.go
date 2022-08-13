package quicksort

import (
	"fmt"
	"testing"
)

func TestSimpleArr(t *testing.T) {
	arr := []int{0, 11, 5, 3, 6, 7, 2, 1, 0, 99}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
