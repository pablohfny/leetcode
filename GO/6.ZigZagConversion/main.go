package main

import "strings"

// func convert(s string, numRows int) string {
// 	length := len(s)

// 	if length == 1 || numRows == 1 || numRows == length {
// 		return s
// 	}

// 	sl := strings.Split(s, "")

// 	result := ""

// 	matrix := make([][]string, numRows)

// 	for i := range matrix {
// 		matrix[i] = make([]string, length)
// 	}

// 	diagColumnAmount := numRows - 2
// 	currIndex := 0
// 	for i := 0; currIndex < length; i++ {
// 		for j := 0; j < numRows && currIndex < length; j++ {
// 			matrix[j][i] = sl[currIndex]
// 			currIndex++
// 		}
// 		diagCounter := 0
// 		for k := numRows - 2; diagCounter < diagColumnAmount && currIndex < length; k-- {
// 			i++
// 			matrix[k][i] = sl[currIndex]
// 			currIndex++
// 			diagCounter++
// 		}
// 	}

// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			val := matrix[i][j]
// 			if val != "" {
// 				result += val
// 			}
// 		}
// 	}

// 	return result
// }

//Simplifications without matrix
func convert(s string, numRows int) string {
	length := len(s)

	if length == 1 || numRows == 1 || numRows == length {
		return s
	}

	sl := strings.Split(s, "")

	result := make([]string, numRows)

	diagColumnAmount := numRows - 2
	currIndex := 0
	for i := 0; currIndex < length; i++ {
		for j := 0; j < numRows && currIndex < length; j++ {
			result[j] += sl[currIndex]
			currIndex++
		}

		diagCounter := 0
		for k := numRows - 2; diagCounter < diagColumnAmount && currIndex < length; k-- {
			i++
			result[k] += sl[currIndex]
			currIndex++
			diagCounter++
		}
	}


	return strings.Join(result, "")
}
