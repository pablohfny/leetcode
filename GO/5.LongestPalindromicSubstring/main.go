package main

import (
	"math"
)

func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	maxPal := ""

	for i, j := 0, length; i < length; i++ {
		pal := s[i:j]
		k := j
		for len(pal) > 0 {
			maxPalLength := len(maxPal)

			if maxPalLength == j-i {
				return maxPal
			}

			if maxPalLength >= k-i {
				break
			}

			if isPalindrome(pal, maxPalLength) {
				maxPal = pal
			}
			k--
			pal = s[i:k]
		}
	}

	return maxPal
}

func isPalindrome(s string, maxLength int) bool {
	if len(s) < maxLength {
		return false
	}

	length := len(s)

	if length == 2 {
		return s[0] == s[1]
	}

	bottomHalf := int(math.Floor(float64(length) / 2.0))
	upperHalf := int(math.Ceil(float64(length) / 2.0))

	for i, j := 0, length-1; i < bottomHalf && j >= upperHalf; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
