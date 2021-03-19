package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{45, 4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	t.Log(values)

	BubbleAsort(values)
	t.Log(values)
}
