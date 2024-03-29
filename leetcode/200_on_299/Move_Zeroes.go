package on299

func MoveZeros(nums []int) {
	if len(nums) == 0 {
		return
	}

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}

			j++
		}
	}
}
