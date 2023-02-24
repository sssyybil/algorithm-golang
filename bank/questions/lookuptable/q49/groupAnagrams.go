package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
*
* 【49. 字母异位词分组】
🌀https://leetcode.cn/problems/group-anagrams/description/
*
* @author  sun
* @date 2022/11/5 11:48
*/

// 时间复杂度为 O(nk logk)，其中 n 是 strs 中字符串的数量，k 是 strs 中字符串的最大长度。需要遍历 n 个字符串，对于每个字符串，都需要 O(klogk) 的时间进行排序以及 O(1) 的时间更新哈希表
// 空间复杂度为 O(nk)
func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	table := make(map[string][]string)

	order := func(s string) string {
		split := strings.Split(s, "")
		sort.Strings(split)
		return strings.Join(split, "")
	}

	for _, v := range strs {
		sortedStr := order(v)
		if s, has := table[sortedStr]; has {
			s = append(s, v)
			table[sortedStr] = s
		} else {
			table[sortedStr] = []string{v}
		}
	}
	for _, v := range table {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
		groupAnagrams([]string{""}),
		groupAnagrams([]string{"a"}),
	)
}
