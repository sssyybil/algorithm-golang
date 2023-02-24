package bottomtoup

import "math"

func MergeSort(nums []int, n int) {

	// 第一层遍历为对进行 merge 的元素个数进行遍历，从第一个元素开始遍历，每次增加两个
	// 也就是说，第一轮归并排序，每次只看一个元素，第二轮看两个元素，第三轮看四个元素，以此类推……
	for sz := 1; sz <= n; sz += sz {

		// 第二层循环为具体每一轮在归并的过程中起始的元素位置，每次 i 的位置的平移为 2 个 size
		// 也就是说，第一轮，将从 [0...size-1]和从 [size...2size-1] 这两部分进行一次归并，第二轮将对 [2size...3size-1] 和 [3size...4size-1] 这两部分进行归并
		for i := 0; i+sz < n; i += sz + sz {

			// 对 arr[i...i+sz-1] 和 arr[i+sz...i+2*sz-1] 进行归并
			merge(nums, i, i+sz-1, int(math.Min(float64(i+sz+sz-1), float64(n-1))))
		}
	}
}

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
