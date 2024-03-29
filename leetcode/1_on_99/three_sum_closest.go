package on99

import (
	"math"
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func ThreeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32

	if n > 2 {
		sort.Ints(nums)

		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[j]

				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}

	return res
}
