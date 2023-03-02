package main

import (
	"math/rand"
	"time"
)

/**
 * 【75. 颜色分类】🐱https://leetcode.cn/problems/sort-colors/description/
 */

// 使用【快速排序】
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	v, j := nums[l], l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[j+1], nums[i] = nums[i], nums[j+1]
			j++
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// sortColorsWithThreeWayQs 使用【三路快速排序】求解
func sortColorsWithThreeWayQs(nums []int) {
	threeWayQs(nums, 0, len(nums)-1)
}

func threeWayQs(nums []int, l, r int) {
	if l >= r {
		return
	}

	// partition
	rand.Seed(time.Now().UnixNano())
	rIndex := rand.Intn(r-l+1) + l
	nums[l], nums[rIndex] = nums[rIndex], nums[l]
	v := nums[l]

	lt, gt, i := l, r+1, l+1
	for i < gt {
		if nums[i] < v {
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			i++
		} else if nums[i] > v {
			nums[gt-1], nums[i] = nums[i], nums[gt-1]
			gt--
		} else {
			// nums[i] == v
			i++
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]

	threeWayQs(nums, l, lt-1)
	threeWayQs(nums, gt, r)
}
