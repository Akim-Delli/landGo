package maxValue

func MaxValue(nums []int) int {
	max := 0

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max

}
