package uptobottom

/**
==> 自底向上的归并排序
*/

func MergeSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	sort(nums, l, mid)
	sort(nums, mid+1, r)
	merge(nums, l, r, mid)
}

// 对 [l...mid] 和 [mid+1...r] 这两部分进行合并
func merge(nums []int, l, r, mid int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}
	// 🌀 i 和 j 指向当前正在考虑的元素，k 指向 i 和 j 相比较后的最终元素应该放入的数组的位置
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = aux[j-l]
			j++
		} else if j > r {
			nums[k] = aux[i-l]
			i++
		} else if aux[i-l] > aux[j-l] {
			nums[k] = aux[j-l]
			j++
		} else {
			nums[k] = aux[i-l]
			i++
		}
	}
}
