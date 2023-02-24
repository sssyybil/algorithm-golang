package quicksort

import (
	"math/rand"
	"time"
)

func DualPivotQuickSort(arr []int) {
	dualPivotSort(arr, 0, len(arr)-1)
}

func dualPivotSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := dualPivotPartition(arr, l, r)
	dualPivotSort(arr, l, p-1)
	dualPivotSort(arr, p+1, r)
}

func dualPivotPartition(arr []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(r-l+1) + l
	arr[l], arr[randomIndex] = arr[randomIndex], arr[l]
	v := arr[l]

	// arr[l...i] <= v; arr[j...r] >= v
	i, j := l+1, r
	for i <= j {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
