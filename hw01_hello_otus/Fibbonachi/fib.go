package fib

func FibN(i int) (int, bool) {
	a0, a1, a2 := 0, 1, 1
	if i == 1 || i == 2 {
		return 1, true
	}
	if i <= 0 {
		return 0, false
	}

	for j := 3; j < i; j++ {
		a0 = a1
		a1 = a2
		a2 = a0 + a1
	}

	return a1 + a2, true
}
