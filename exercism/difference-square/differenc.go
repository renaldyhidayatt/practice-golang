package differencesquare

func SquareOfSum(n int) int {
	result := (n * (n + 1)) / 2
	return result * result
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
