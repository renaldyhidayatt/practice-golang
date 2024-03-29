package math

func Phi(n int64) int64 {
	result := n

	for i := int64(2); i*i <= n; i += 1 {
		if n%i == 0 {
			for {
				if n%i != 0 {
					break
				}
				n /= i
			}

			result -= result / i
		}
	}

	if n > 1 {
		result -= result / n
	}

	return result
}
