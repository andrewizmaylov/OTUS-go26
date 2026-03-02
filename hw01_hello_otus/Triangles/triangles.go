package triangles

const (
	Impossible = iota
	Acute
	Right
	Obtuse
)

func findMax(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}

	return c
}

func findC1(a, b, c, maximum int) int {
	if a != maximum {
		return a
	} else if b != maximum {
		return b
	}

	return c
}

func findC2(a, b, c, maximum, c1 int) int {
	if a != maximum && a != c1 {
		return a
	} else if b != maximum && b != c1 {
		return b
	}

	return c
}

func checkTriangle(a, b, c int) int {
	if a + b <= c || b + c <= a || c + a <= b {
		return Impossible
	}

	maximum := findMax(a, b, c)

	c1 := findC1(a, b, c, maximum)
	c2 := findC2(a, b, c, maximum, c1)

	if (c1*c1 + c2*c2) == maximum*maximum {
		return Right
	} else if (c1*c1 + c2*c2) > maximum*maximum {
		return Obtuse
	} else if (c1*c1 + c2*c2) < maximum*maximum {
		return Acute
	}

	return Impossible
}
