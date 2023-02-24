package uptobottom

/**
==> 自底向上的归并排序
==> 优化 A：归并过程保证了从 [l...mid] 和从 [mid+1...r] 都是有序的，因此只有当 nums[mid] > nums[mid+1] 时，才需要执行合并的操作
*/

func SortOptimizeA(nums []int) {
	sortA(nums, 0, len(nums)-1)
}

func sortA(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	sortA(nums, l, mid)
	sortA(nums, mid+1, r)

	// 🍑 优化点
	if nums[mid] > nums[mid+1] {
		merge(nums, l, r, mid)
	}
}
