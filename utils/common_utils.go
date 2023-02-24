package utils

import (
	"sort"
)

func IsSliceEquals(nums1, nums2 []int) bool {
	n1, n2 := len(nums1), len(nums2)
	if n1 != n2 {
		return false
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 0; i < n1; i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

// TODO 判断两个二维数组是否相等
//func IsTwoDimensionalArrayEqual(nums1, nums2 [][]int) bool {
//	n1, n2 := len(nums1), len(nums2)
//	if n1 != n2 {
//		return false
//	}
//
//	arrayToMap := func(array [][]int) map[string]int {
//		record := make(map[string]int)
//		for _, x := range array {
//			sort.Ints(x)
//			var s string
//			strings.Trim(strings.Join(strings.Fields(fmt.Sprint(x)), s), "[]")
//			if v, has := record[s]; has {
//				record[s] = v + 1
//			} else {
//				record[s] = 1
//			}
//		}
//		return record
//	}
//
//	m1, m2 := arrayToMap(nums1), arrayToMap(nums2)
//	for k, v := range m1 {
//		if value, has := m2[k]; has && v != 0 {
//			m2[k] = value - 1
//		} else {
//			return false
//		}
//	}
//	return true
//}

func IsMatrixEqual(matrix1, matrix2 [][]int) bool {
	n1, n2 := len(matrix1), len(matrix2)
	if n1 != n2 {
		return false
	}

	row := 0
	for row < n1 {
		// 若列的长度不等，则判定两个二维数组不同
		if len(matrix1[row]) != len(matrix2[row]) {
			return false
		}
		column := 0
		for column < len(matrix1[row]) {
			if matrix1[row][column] != matrix2[row][column] {
				return false
			}
			column++
		}
		row++
	}
	return true
}
