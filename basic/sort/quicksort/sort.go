package quicksort

import (
	"algorithm-golang/basic/sort/insertsort"
	"math/rand"
	"time"
)

func Sort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	// 使用【插入排序】来优化算法
	if r-l <= 15 {
		insertsort.SortWithCopy(arr, l, r)
	}

	p := dualPivotPartition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// partition => 对 arr[l..r] 部分进行 partition 操作
// 返回 p，使得 arr[l...p-1] < arr[p]；arr[p+1...r] > arr[p]
func partition(arr []int, l, r int) int {
	// 🍑【优化快速排序】
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(r-l+1) + l
	arr[l], arr[randomIndex] = arr[randomIndex], arr[l]
	v := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
