func topThreeNaive(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	max1 := math.MinInt64
	for _, val := range arr {
		if val > max1 {
			max1 = val
		}
	}

	max2 := math.MinInt64
	for _, val := range arr {
		if val > max2 && val != max1 {
			max2 = val
		}
	}

	max3 := math.MinInt64
	for _, val := range arr {
		if val > max3 && val != max1 && val != max2 {
			max3 = val
		}
	}

	result := []int{}
	if max1 != math.MinInt64 {
		result = append(result, max1)
	}
	if max2 != math.MinInt64 {
		result = append(result, max2)
	}
	if max3 != math.MinInt64 {
		result = append(result, max3)
	}
	return result
}

