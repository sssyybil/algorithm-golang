package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
*
* 【18. 四数之和】
🌀https://leetcode.cn/problems/4sum/description/
*
* @author  sun
* @date 2022/11/3 15:11
*/

// 三层嵌套遍历，时间复杂度为 O(n^3)；因使用到了额外的空间去重，因此空间复杂度为 O(N)
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	var ans [][]int
	if n < 4 {
		return ans
	}
	sort.Ints(nums)
	// 🤔开始考虑的是，按升序排序后的数组，如果最小值也比 target 要大，那么四个数相加之后肯定也会比 target 要大。
	// 实际这种只符合 nums 内所有值均为正数的情况，如果排序后的数组前几位均为负数，那么即使符合上述条件，四个数的和也是与 target 值相同的
	//if nums[0] > target {
	//	return ans
	//}
	if nums[0] > 0 && nums[0] > target {
		return ans
	}
	duplicate := make(map[string]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])
				if sum == target && !duplicate[key] {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					duplicate[key] = true
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(
		fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11), // [[-5,-4,-3,1]]
		fourSum([]int{1, 0, -1, 0, -2, 2}, 0),           // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
		fourSum([]int{2, 2, 2, 2, 2}, 8),                // [[2,2,2,2]]
	)
}
