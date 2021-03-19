package sort

func maxHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		// 判断是否存在child节点
		if child > end {
			break
		}
		// 判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] < c[child+1] {
			child += 1
		}
		if c[root] < c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		// 判断是否存在child节点
		if child > end {
			break
		}
		// 判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

//在c数组中找出num个最大值
func HeapSort(c []int, num int) []int {
	m := len(c) - 1
	createHeap(c[:num], num-1)

	for i := num; i <= m; i++ {
		if c[0] < c[i] {
			c[0], c[i] = c[i], c[0]
			createHeap(c[:num], num-1)
		}
	}

	return c
}
func createHeap(arr []int, end int) {
	for start := end / 2; start >= 0; start-- {
		maxHeap(start, end, arr)
	}
}
