package insertsort

func Sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				// 前面的元素均为有序状态，因此当符合 nums[j] >= nums[j-1] 时，表明无需再向前遍历了
				break
			}
		}
	}
}

func SortOptimize(nums []int) {
	for i := 1; i < len(nums); i++ {
		copyI, j := nums[i], i
		for ; j > 0 && nums[j-1] > copyI; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = copyI
	}
}

func SortWithCopySimpler(nums []int) {
	SortWithCopy(nums, 0, len(nums))
}

func SortWithCopy(nums []int, l, r int) {
	for i := l; i < r; i++ {
		copyI, j := nums[i], i
		for ; j > 0 && nums[j-1] > copyI; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = copyI
	}
}
