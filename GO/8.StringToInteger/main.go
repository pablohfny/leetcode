package main

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	s = strings.Trim(s, " ")
	res := 0
	sign := 1
	started := false
	for _, char := range s {
		if !started && (char == '-' || char == '+') {
			if char == '-' {
				sign = -1
			}
			started = true
			continue
		}
		if !unicode.IsDigit(char) {
			break
		}
		started = true
		res = res*10 + int(char-'0')

		if res*sign < math.MinInt32 {
			return math.MinInt32
		}
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return res * sign
}
