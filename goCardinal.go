package goCardinal

import (
	"strings"
)

// Slices for units and tens
var units = []string{
	"", "one", "two", "three", "four", "five", "six", "seven",
	"eight", "nine", "ten", "eleven", "twelve", "thirteen",
	"fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var unitsOrdinals = []string{
	"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "tenth", "eleventh", "twelfth", "thirteenth",
	"fourteenth", "fifteenth", "sixteenth", "seventeenth", "eighteenth", "nineteenth",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty",
	"seventy", "eighty", "ninety",
}

var tensOrdinals = []string{
	"", "", "twentieth", "thirtieth", "fortieth", "fiftieth", "sixtieth",
	"seventieth", "eightieth", "ninetieth",
}

var magnitudes = []struct {
	value           int64
	singular        string
	singularOrdinal string
}{
	{1e18, "quintillion", "quintillionth"},
	{1e15, "quadrillion", "quadrillionth"},
	{1e12, "trillion", "trillionth"},
	{1e9, "billion", "billionth"},
	{1e6, "million", "millionth"},
	{1e3, "thousand", "thousandth"},
	{1e2, "hundred", "hundredth"},
}

// NumberToOrdinal converts an integer to its ordinal word form.
// Supports numbers up to 9,223,372,036,854,775,807 (max int64 value).
func NumberToOrdinal(n int64) string {
	if n == 0 {
		return "Zeroth"
	}

	var words []string
	remainder := n
	ordinalAlreadySet := false // Flag to indicate if ordinal conversion has been done

	// Process magnitudes (quintillion, quadrillion, etc.)
	for _, magnitude := range magnitudes {
		if remainder >= magnitude.value {
			count := remainder / magnitude.value
			remainder %= magnitude.value

			// Get the word for the count (e.g., "one", "two")
			var countWord string
			if count < 20 {
				countWord = units[count]
			} else if count < 100 {
				if count%10 == 0 {
					countWord = tens[count/10]
				} else {
					countWord = tens[count/10] + "-" + units[count%10]
				}
			} else {
				// For counts >= 100, recursively process
				countWord = numberToWords(count)
			}

			// Append the magnitude word
			words = append(words, countWord, magnitude.singular)

			if remainder == 0 {
				// Convert magnitude to its ordinal form
				words[len(words)-1] = magnitude.singularOrdinal
				ordinalAlreadySet = true // Set the flag
				break
			}
		}
	}

	// Process tens and units
	if remainder > 0 {
		if remainder < 20 {
			// For numbers less than 20
			words = append(words, unitsOrdinals[remainder])
		} else if remainder%10 == 0 {
			// Exact tens (e.g., 20, 30)
			words = append(words, tensOrdinals[remainder/10])
		} else {
			// Numbers between tens (e.g., 21, 47)
			words = append(words, tens[remainder/10]+"-"+unitsOrdinals[remainder%10])
		}
	} else if len(words) > 0 && !ordinalAlreadySet {
		// If remainder is zero and words exist, and ordinal not already set
		lastIndex := len(words) - 1
		lastWord := words[lastIndex]
		if strings.HasSuffix(lastWord, "y") {
			words[lastIndex] = lastWord[:len(lastWord)-1] + "ieth"
		} else {
			words[lastIndex] = lastWord + "th"
		}
	}

	// Capitalize the first letter
	result := strings.Join(words, " ")
	if len(result) > 0 {
		result = strings.ToUpper(result[:1]) + result[1:]
	}

	return result
}

// Helper function to convert numbers to words without ordinal suffixes.
func numberToWords(n int64) string {
	var words []string
	remainder := n

	// Process magnitudes (quintillion, quadrillion, etc.)
	for _, magnitude := range magnitudes {
		if remainder >= magnitude.value {
			count := remainder / magnitude.value
			remainder %= magnitude.value

			// Get the word for the count
			var countWord string
			if count < 20 {
				countWord = units[count]
			} else if count < 100 {
				if count%10 == 0 {
					countWord = tens[count/10]
				} else {
					countWord = tens[count/10] + "-" + units[count%10]
				}
			} else {
				// For counts >= 100, recursively process
				countWord = numberToWords(count)
			}

			words = append(words, countWord, magnitude.singular)
		}
	}

	// Process tens and units
	if remainder > 0 {
		if remainder < 20 {
			words = append(words, units[remainder])
		} else if remainder%10 == 0 {
			words = append(words, tens[remainder/10])
		} else {
			words = append(words, tens[remainder/10]+"-"+units[remainder%10])
		}
	}

	return strings.Join(words, " ")
}
