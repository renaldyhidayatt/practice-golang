package gcd

func Iterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

