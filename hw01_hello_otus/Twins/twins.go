package twins

func CheckTwins(i int) bool {
	if i <= 1 {
		return false
	}
	if i == 2 || i == 3 {
		return true
	}
	for j := 2; j < i; j ++ {
		if j < i && i % j == 0 {
			return false
		}
	}

	return true
}
