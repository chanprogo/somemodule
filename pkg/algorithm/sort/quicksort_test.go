package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {

	arr := []int64{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	t.Log(arr)

	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
