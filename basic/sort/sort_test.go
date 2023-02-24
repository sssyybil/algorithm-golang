package main

import (
	"algorithm-golang/basic/sort/helper"
	"algorithm-golang/basic/sort/mergesort/uptobottom"
	"testing"
)

var nearlySortedNums = helper.GenerateNearlyOrderArray(5000000, 10)

func TestMergeSort(t *testing.T) {
	var nums = helper.GenerateRandomArray(1000000, 9, 9999)
	uptobottom.MergeSort(nums)
}

func TestMergeSortOptimizeA(t *testing.T) {
	numsCopy := make([]int, len(nearlySortedNums))
	copy(numsCopy, numsCopy)

	uptobottom.SortOptimizeA(numsCopy)
}

func TestMergeSortOptimizeB(t *testing.T) {
	numsCopy := make([]int, len(nearlySortedNums))
	copy(numsCopy, numsCopy)

	uptobottom.SortOptimizeB(numsCopy)
}
