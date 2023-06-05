package usecases

import (
	"strings"
)

// Static chars definition
var vowels = []string{"a","e","i","o","u"}
var ignoredChars = []string{".", " "}

func CalculateSuitabilityScore(driverName string, oddOrEven string, addressName string) float64 {
	// Ignoring . and space characters
	for _, char := range ignoredChars {
		driverName = strings.ReplaceAll(driverName, char, "")
		addressName = strings.ReplaceAll(addressName, char, "")
	}
	driverName = strings.ToLower(driverName)
	// Counting vowels, maybe can try with runes but using strings library for now
	total_vowel_count := 0
	for _, vowel := range vowels {
		total_vowel_count += strings.Count(driverName, vowel)
	}
	// For consonant count just total count for drivername minus the vowel count
	total_consonant_count := len(driverName) - total_vowel_count
	//fmt.Println(total_vowel_count, total_consonant_count)
	result := 0.0
	// For vowels if odd 1.5x
	if oddOrEven == "odd" {
		result = 1.5 * float64(total_vowel_count)
	}
	// For consonant added 1x, but if someone need to configure the model just by adjusting the number
	if oddOrEven == "even" {
		result = 1 * float64(total_consonant_count)
	}
	// Multiply by 2.5 if Common Factors are more than 0
	if len(GetCommonFactors(len(addressName), len(driverName))) > 0 {
		result = 2.5 * result
	}
	return result
}
