package main

import "fmt"

/**
 * 【5. 最长回文子串】🐱https://leetcode.cn/problems/longest-palindromic-substring/
 */

// TODO repeat

// 找出所有的子串 + 判断其是否是回文
func longestPalindrome(s string) string {
	memory := make(map[string]int)
	n, maxLen, res := len(s), 0, string(s[0])
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			subStr := s[i : j+1]
			// 将子串存入 memory 中，防止重复判断
			_, ok := memory[subStr]
			if !ok && isPalindrome(subStr) {
				memory[subStr] = 1
				if len(subStr) > maxLen {
					maxLen = len(subStr)
					res = subStr
				}
			}
		}
	}
	//fmt.Printf("len(memory)=%d, memory=%v", len(memory), memory)
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

// 使用【动态规划】解决
func longestPalindromeWithDp(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen, begin := 1, 0
	// dp[i][j] 表示 s[i...j] 是否是回文串
	dp := make([][]bool, n)
	// 初始化所有长度为 1 的子串都是回文串
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := l + i - 1
			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func main() {
	fmt.Println(
		longestPalindromeWithDp("babad"),
		//longestPalindrome("vvgogaewginhopuxzwyryobjjpieyhwfopiowxmyylvcgsnhfcnogpqpukzmnpewavoutbloyrrhatimmxfqmcwgfebuoqbbgvubbkjfvxivjfijlpvtsnhagzfptahhyojwzamayoiegkenycnkxzkambimhdykdcxyyfjnvztzypmfczdhhnkmfuvgkhzbwmjznykkagqdrueohgcmeidjqsvfugcioeduohprjtfdmtzosnhoiganffarokxjifzzxhixdzycwfheqqegduzwtiacmdhqfmxhazcxsqyrvrihfqzjxvawdeandnwgjlquvzadruiqmcsgibglhicsvzqknztqpkiihdoisxipkourentfvrltieihiktgzswtgcmmlfrjifqinhrbplbsgswqlbjkyxjwoshsvxlhlpgzwbmxzwaemtowcxwourjwmmwjruowxcwotmeawzxmbwzgplhlxvshsowjxykjblqwsgsblpbrhniqfijrflmmcgtwszgtkihieitlrvftneruokpixsiodhiikpqtznkqzvscihlgbigscmqiurdazvuqljgwndnaedwavxjzqfhirvryqsxczahxmfqhdmcaitwzudgeqqehfwcyzdxihxzzfijxkoraffnagiohnsoztmdftjrphoudeoicgufvsqjdiemcghoeurdqgakkynzjmwbzhkgvufmknhhdzcfmpyztzvnjfyyxcdkydhmibmakzxkncynekgeioyamazwjoyhhatpfzgahnstvpljifjvixvfjkbbuvgbbqoubefgwcmqfxmmitahrryolbtuovawepnmzkupqpgoncfhnsgcvlyymxwoipofwhyeipjjboyrywzxupohnigweagogvv"),
	)

}
