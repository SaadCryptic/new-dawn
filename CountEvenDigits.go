package main

func CountEvenDigits(digits string) int {
	count := 0
	for _, char := range digits {
		if char >= '0' && char <= '9' {
			if (char-'0')%2 == 0 {
				count++
			}
		}

	}
	return count
}
