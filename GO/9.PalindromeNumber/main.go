package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	tail := 0

	for x > tail {
		tail = tail*10 + x%10
		x /= 10
	}

	return x == tail || x == (tail/10)
}
