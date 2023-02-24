package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

/**
 * 【461. 汉明距离】🐱https://leetcode.cn/problems/hamming-distance/
 */

// hammingDistance 使用内置函数将 x 和 y 转换成二进制后，再遍历两个二进制字符，计算其中的不相同的字符个数
func hammingDistance(x int, y int) int {
	binaryX := strconv.FormatInt(int64(x), 2)
	binaryY := strconv.FormatInt(int64(y), 2)

	nX, nY := len(binaryX), len(binaryY)
	if nX != nY {
		maxLen := int(math.Max(float64(nX), float64(nY)))
		if nX < nY {
			for len(binaryX) < maxLen {
				binaryX = "0" + binaryX
			}
		} else {
			for len(binaryY) < maxLen {
				binaryY = "0" + binaryY
			}
		}
	}

	fmt.Printf("x = %v , y = %v\n", binaryX, binaryY)

	res := 0
	for i := 0; i < len(binaryX); i++ {
		if binaryX[i] != binaryY[i] {
			res++
		}
	}

	return res
}

// hammingDistanceWithBuildInFun 使用内置的计算二进制数字中 1 的个数的函数。推荐在工程中使用此方法。
func hammingDistanceWithBuildInFun(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println(
		hammingDistance(680142203, 1111953568),
		hammingDistance(3, 1),
	)
}
