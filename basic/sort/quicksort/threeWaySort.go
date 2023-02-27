package quicksort

import (
	"math/rand"
	"time"
)

/**
 * 【三路快速排序法】
 * 将 arr[l...r] 分为 <v; ==v; >v 三部分
 * 之后递归对 <v 和 >v 两部分继续进行三路快速排序
 */

func ThreeWayQuickSort(arr []int) {
	threeWaySort(arr, 0, len(arr)-1)
}

func threeWaySort(arr []int, l, r int) {
	if l >= r {
		return
	}

	// partition 过程，三路快排与其它两种快排方式不同，在每次的排序过程中，arr[l...lt] 与 arr[gt...r] 都需要进行 partition 过程
	// 因此将 partition 过程合并到当前函数中
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(r-l+1) + l
	arr[l], arr[randomIndex] = arr[randomIndex], arr[l]
	v := arr[l]

	// arr[l+1...lt] < v; arr[gt...r] > v; arr[lt+1...i) == v
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < v {
			arr[lt+1], arr[i] = arr[i], arr[lt+1]
			lt++
			i++
		} else if arr[i] > v {
			arr[gt-1], arr[i] = arr[i], arr[gt-1]
			// 此处将当前待处理元素 arr[i] 与 arr[gt-1] 元素交换后，arr[i] 重新指向的元素仍然是待处理的，因此无需执行 i++
			gt--
		} else { // arr[i] == v
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]

	threeWaySort(arr, l, lt-1)
	threeWaySort(arr, gt, r)
}
