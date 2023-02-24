package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
*
* 【15. 三数之和】
🌀https://leetcode.cn/problems/3sum/
*
* @author  sun
* @date 2022/11/1 10:13
*/

// 三层嵌套遍历，时间复杂度为 O(n^3) ==> 超时……
func threeSumA(nums []int) [][]int {
	var ans [][]int
	duplicate := make(map[string]bool)
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {

					cur := []int{nums[i], nums[j], nums[k]}
					sort.Ints(cur)
					str := numsToString(cur)

					if !duplicate[str] {
						ans = append(ans, cur)
						duplicate[str] = true
					}
				}
			}
		}
	}

	return ans
}

// 两层嵌套遍历，第二层使用双指针方式
func threeSumB(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n <= 2 {
		return ans
	}

	sort.Ints(nums)
	duplicate := make(map[string]bool)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			curSum := nums[i] + nums[l] + nums[r]
			if curSum == 0 {
				cur := []int{nums[i], nums[l], nums[r]}
				sort.Ints(cur)
				str := numsToString(cur)

				if !duplicate[str] {
					ans = append(ans, cur)
					duplicate[str] = true
				}
				l++
				r--
			} else if curSum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

func numsToString(nums []int) string {
	ans := ""
	for _, v := range nums {
		ans += strconv.Itoa(v)
	}
	return ans
}

func threeSumC(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n <= 2 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < n; i++ {
		// 对于已经排好序的数组，如果 nums[i] > 0，则表示 i 后面的元素之和不可能等于 0，没有再继续遍历的必要
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			curSum := nums[i] + nums[l] + nums[r]
			if curSum == 0 {
				fmt.Printf("i=%v, l=%v, r=%v\n", i, l, r)
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if curSum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(
		threeSumA([]int{-1, 0, 1, 2, -1, -4}), // [[-1,-1,2],[-1,0,1]]
		threeSumA([]int{0, 1, 1}),             //[]
		threeSumA([]int{0, 0, 0}),             // [[0,0,0]]
	)

	fmt.Println(
		threeSumB([]int{-1, 0, 1, 2, -1, -4}), // [[-1,-1,2],[-1,0,1]]
		threeSumB([]int{0, 1, 1}),             //[]
		threeSumB([]int{0, 0, 0}),             // [[0,0,0]]
	)

	fmt.Println(
		threeSumC([]int{-2, 0, 0, 2, 2}),      // [[-2,0,2]]
		threeSumC([]int{-1, 0, 1, 2, -1, -4}), // [[-1,-1,2],[-1,0,1]]
		threeSumC([]int{0, 1, 1}),             //[]
		threeSumC([]int{0, 0, 0}),             // [[0,0,0]]
	)

}
