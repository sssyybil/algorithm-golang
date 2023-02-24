package main

func firstUniqChar_A(s string) byte {
	tmp := make(map[rune]int)
	for _, r := range []rune(s) {
		if count, has := tmp[r]; has {
			tmp[r] = count + 2
		} else {
			tmp[r] = 1
		}
	}
	for _, r := range []rune(s) {
		if tmp[r] == 1 {
			return byte(r)
		}
	}
	return ' '
}

func firstUniqChar_B(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
