package on299

import "math"

func NumSquares(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

func isPerfectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}

func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}
