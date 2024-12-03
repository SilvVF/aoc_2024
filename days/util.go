package days

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func absDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}

func removeCpy(slice []int, s int) []int {
	new := make([]int, len(slice)-1)
	idx := 0
	for i, n := range slice {
		if i == s {
			continue
		}
		new[idx] = n
		idx += 1
	}
	return new
}
