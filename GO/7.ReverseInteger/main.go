package main

import "math"

//String conversion
// func reverse(x int) int {
// 	s := ""

// 	if x < 0 {
// 		s = strconv.Itoa(-x)
// 	} else {
// 		s = strconv.Itoa(x)
// 	}

// 	sAr := strings.Split(s, "")
// 	sResult := ""

// 	for i := len(sAr) - 1; i >= 0; i-- {
// 		sResult += sAr[i]
// 	}

// 	val, _ := strconv.ParseInt(sResult, 10, 0)

// 	if val > math.MaxInt32 {
// 		return 0
// 	}

// 	if x < 0 {
// 		return -int(val)
// 	}

// 	return int(val)
// }

// Mathematical conversion
func reverse(x int) int {
	reverse := 0
	for {
		reverse = reverse*10 + (x % 10)
		x = x / 10
		if x == 0 {
			break
		}
	}
	if reverse < math.MinInt32 || reverse > math.MaxInt32 {
		return 0
	}
	return reverse
}
