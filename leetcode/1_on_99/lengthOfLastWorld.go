package on99


func LengthOfLastWord(s string) int{
	last := len(s) - 1

	for last >= 0 && s[last] == ' ' {
		last--
	}

	if last < 0{
		return 0

	}
	first := last

	for first >= 0 && s[first] != ' '{
		first--
	}

	return last - first
}