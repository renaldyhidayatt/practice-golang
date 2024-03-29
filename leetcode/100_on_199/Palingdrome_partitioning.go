package on199

func Partition131(s string) [][]string {
	result := [][]string{}
	size := len(s)

	if size == 0 {
		return result
	}

	current := make([]string, 0, size)
	dfs131(s, 0, current, &result)

	return result
}

func dfs131(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)

	if start == end {
		temp := make([]string, len(cur))

		copy(temp, cur)
		*result = append(*result, temp)

		return
	}

	for i := start; i < end; i++ {
		if isPal(s, start, i) {
			dfs131(s, i+1, append(cur, s[start:i+1]), result)
		}
	}
}

func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
