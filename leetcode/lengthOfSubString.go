package main

func LengthOfLongestSubstring(s string) int {
	charMap := make([]int, 256)
	for i := range charMap {
		charMap[i] = -1
	}
	i, maxLen := 0, 0
	for j := 0; j < len(s); j++ {
		if charMap[s[j]] >= i {
			i = charMap[s[j]] + 1
		}
		charMap[s[j]] = j
		maxLen = Max(j-i+1, maxLen)
	}
	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}