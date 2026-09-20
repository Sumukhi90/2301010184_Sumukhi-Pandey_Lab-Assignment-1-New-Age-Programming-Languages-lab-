package custompkg

import (
	"strings"
)

// Reverse reverses a string
func Reverse(text string) string {
	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// ToUpper converts string to uppercase
func ToUpper(text string) string {
	return strings.ToUpper(text)
}

// Factorial calculates factorial of a number
func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

// Power calculates the power of a number
func Power(base, exponent float64) float64 {
	result := 1.0

	for i := 0; i < int(exponent); i++ {
		result *= base
	}

	return result
}
