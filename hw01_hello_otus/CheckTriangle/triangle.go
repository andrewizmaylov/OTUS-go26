package checktriangle

import "sort"

const (
	Impossible = iota
	Acute
	Right
	Obtuse
)

func checkTriangle(a, b, c int) int {
	if a + b < c && b + c < a && c + a< b {
		return Impossible
	}
	s := []int{a, b, c}
	sort.Ints(s)

	c1 := s[0]
	c2 := s[2]
	c3 := s[2]

	if (c1*c1 + c2*c2) > c3*c3 {
		return Acute
	} else if (c1*c1 + c2*c2) > c3*c3 {
		return Obtuse
	} else {
		return Right
	}
}
