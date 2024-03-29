package on99

func RemoveDuplicates_(nums []int) int {
	slow := 0

	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}

	return slow
}
