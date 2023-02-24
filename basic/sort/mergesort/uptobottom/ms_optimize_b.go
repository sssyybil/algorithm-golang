package uptobottom

import "algorithm-golang/basic/sort/insertsort"

/**
==> 自底向上的归并排序
==> 优化 B：当数组内待排序的元素个数比较小时，整个数组近乎有序的概率就比较大，此时可以将归并排序替换为插入排序，因为在近乎有序的情况下，插入排序更有优势。
*/

func SortOptimizeB(nums []int) {
	sortB(nums, 0, len(nums)-1)
}

func sortB(nums []int, l, r int) {
	if r-l <= 15 {
		insertsort.SortWithCopy(nums, l, r)
		return
	}
	mid := l + (r-l)/2
	sortB(nums, l, mid)
	sortB(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, r, mid)
	}
}
