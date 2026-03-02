package palindrome

import (
	"unicode"
	"unicode/utf8"
)

func InspectPhrase(s string) (char rune, count int, isComplex bool) {
	var cache = make(map[rune]int, 255)
	space := []rune(" ")
	str := []rune(s)
	length := utf8.RuneCountInString(s) - 1

	for i := 0; i < length; i ++ {
		if str[i] != space[0] {
			key := unicode.ToLower(str[i])
			cache[key] ++
		}
	}

	char = ' '
	count = 0

	for v, c := range cache {
		if count < c {
			count = c
			char = v
		}
	}

	isComplex = (float32(count) / float32(length)) > 0.25
	return char, count, isComplex
}

func IsPalindromeFromString(s string) bool {
	str := []rune(s)
	space := []rune(" ")

	s1 := 0
	s2 := utf8.RuneCountInString(s) - 1

	for s1 < s2 {
		if str[s1] == space[0] {
			s1 ++
			continue
		}
		if str[s2] == space[0] {
			s2 --
			continue
		}
		if str[s1] != str[s2] {
			return false
		}
		s1 ++
		s2 --

	}

	return true
}

func IsPalindrome(s string) bool {
	p1 := 0
	p2 := utf8.RuneCountInString(s) - 1
	// переобразование строки в слайс рун - аналог php str_split('adnfvdf')
	q := []rune(s)

	for p1 < p2 {
		if q[p1] != q[p2] {
			return false
		}
		p1 ++
		p2 --
	}

	return true
}

// func converter() {
// 	str := "Немного текста"

// 	runeSlice := []rune(str)
// 	s = string(runeSlice)

// 	byteSlice := []byte(str)
// 	s = string(byteSlice)

// }
