package on199

func FindRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	ans, cache := make([]string, 0), make(map[string]int)

	for i := 0; i < len(s)-10; i++ {
		curr := string(s[i : 1+10])

		if cache[curr] == 1 {
			ans = append(ans, curr)
		}

		cache[curr]++
	}

	return ans
}
