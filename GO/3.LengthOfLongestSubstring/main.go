package main

// import "strings"

// func lengthOfLongestSubstring(s string) int {
// 	total := len(s)

// 	if total < 2 {
// 		return total
// 	}

// 	allCharsMap := make(map[string]int)
// 	allCharsTotal := 0
// 	sl := strings.Split(s, "")

// 	for _, char := range sl {
// 		if _, ok := allCharsMap[char]; ok {
// 			continue
// 		}
// 		allCharsTotal += 1
// 		allCharsMap[char] = 1
// 	}

// 	longest := ""
// 	aux := longest

// 	charMap := make(map[string]int)
// 	longestLen := len(longest)

// 	for i := 0; i < len(sl); i++ {
// 		char := sl[i]
// 		longestLen = len(longest)

// 		if index, ok := charMap[char]; ok {
// 			i = index
// 			aux = ""
// 			charMap = make(map[string]int)
// 			continue
// 		} else {
// 			aux += char
// 			charMap[char] = i
// 		}

// 		auxLen := len(aux)

// 		if auxLen > longestLen {
// 			longest = aux
// 			continue
// 		}

// 		if longestLen == allCharsTotal || (auxLen+(len(sl)-(i+1))) <= longestLen {
// 			return longestLen
// 		}
// 	}

// 	return len(longest)
// }

func lengthOfLongestSubstring(input string) int {
	var res, left, right int

	encounteredChars := make(map[byte]int)

	for left, right = 0, 0; right < len(input); right++ {
		if prevCharIndex, stopExpandWindow := encounteredChars[input[right]]; stopExpandWindow && prevCharIndex >= left {
			res = max(res, right-left)
			left = prevCharIndex + 1
		}

		encounteredChars[input[right]] = right
	}

	res = max(res, right-left)
	return res
}
