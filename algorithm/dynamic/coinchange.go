package dynamic

func CoinChange(coins []int32, amount int32) int32 {
	combination := make([]int32, amount)

	combination[0] = 1

	for _, c := range coins {
		for i := c; i < amount; i++ {
			combination[i] += combination[i-c]
		}
	}

	return combination[amount-1]
}
