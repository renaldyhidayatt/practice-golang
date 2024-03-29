package search

func selectK(array []int, l, r, idx int) int {
	index := partition(array, l, r)

	if index == idx {
		return array[index]
	}

	if index < idx {
		return selectK(array, index+1, r, idx)
	}
	return selectK(array, l, index, idx)
}

func partition(array []int, l, r int) int {
	elem, j := array[l], l+1

	for i := l + 1; i < r; i++ {
		if array[i] <= elem {
			array[i], array[j] = array[j], array[i]

			j++
		}
	}

	array[l], array[j-1] = array[j-1], array[l]

	return j - 1
}
