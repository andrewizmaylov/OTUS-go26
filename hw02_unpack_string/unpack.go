package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (unpacked string, err error) {
	runes := []rune(s)
	if len(runes) == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var sb strings.Builder
	var prev rune
	var escapeMode bool

	for i := range len(runes) {
		current := runes[i]

		// Handle escape sequences
		if current == '\\' && !escapeMode {
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			escapeMode = true
			if prev != 0 && !unicode.IsDigit(prev) {
				sb.WriteRune(prev)
			}
			prev = runes[i+1]
			continue
		}
		// Escaped backslash: set prev to literal, don't write yet (next char might be digit)
		if escapeMode && current == '\\' {
			prev = '\\'
			escapeMode = false
			continue
		}
		// If current is digit and not in escape mode
		if unicode.IsDigit(current) && !escapeMode {
			// Consecutive digits in source (e.g. "10" in "aaa10b") are invalid.
			// Escaped digit followed by digit (e.g. \45) is valid: literal 4, then repeat 5 times.
			prevRuneInSourceWasDigit := i > 0 && unicode.IsDigit(runes[i-1])
			prevRuneWasEscaped := i >= 2 && runes[i-2] == '\\'
			if prevRuneInSourceWasDigit && !prevRuneWasEscaped {
				return "", ErrInvalidString
			}

			// Convert digit to number
			count, _ := strconv.Atoi(string(current))

			if count == 0 {
				// Remove previous character (don't write it)
				// Since we can't remove from builder, we need to handle differently
				// For now, we'll just skip writing it
				prev = 0
			} else {
				// Write previous character count times
				if unicode.IsDigit(prev) {
					count--
				}
				for range count {
					sb.WriteRune(prev)
				}
			}
			prev = current
		} else {
			// Write previous character: either non-digit, or escaped literal (digit/backslash)
			if prev != 0 && (!unicode.IsDigit(prev) || escapeMode) {
				sb.WriteRune(prev)
			}
			prev = current
			escapeMode = false
		}
	}

	if prev != 0 && !unicode.IsDigit(prev) {
		sb.WriteRune(prev)
	}

	result := sb.String()

	return result, nil
}
