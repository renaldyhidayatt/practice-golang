package binary

import "math"

func FastInverseSqrt(number float32) float32 {
	var i uint32
	var y, x2 float32
	const threehalfs float32 = 1.5

	x2 = number * float32(0.5)
	y = number
	i = math.Float32bits(y)
	i = 0x5f3759df - (i >> 1)
	y = math.Float32frombits(i)

	y = y * (threehalfs - (x2 * y * y))
	y = y * (threehalfs - (x2 * y * y))
	return y
}
