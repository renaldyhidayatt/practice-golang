package on399

func lastRemaining(n int) int {
	start, dir, step := 1, true, 1

	for n > 1 {
		if dir {
			start += step
		} else {
			if n%2 == 1 {
				start += step
			}
		}

		dir = !dir

		n >>= 1

		step <<= 1
	}

	return start
}
