package goCardinal

import (
	"math"
	"strconv"
	"testing"
)

func TestNumberToOrdinal(t *testing.T) {
	cases := []struct {
		input int64
		want  string
	}{
		{0, "Zeroth"},
		{1, "First"},
		{21, "Twenty-first"},
		{100, "One hundredth"},
		{342, "Three hundred forty-second"},
		{-1, "Minus first"},
		{-21, "Minus twenty-first"},
		{-100, "Minus one hundredth"},
		{-342, "Minus three hundred forty-second"},
		{-1000000000000000000, "Minus one quintillionth"},
		{math.MaxInt64, "Nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seventh"},
		{math.MinInt64, "Minus nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred eighth"},
	}
	for _, tc := range cases {
		t.Run(strconv.FormatInt(tc.input, 10), func(t *testing.T) {
			if got := NumberToOrdinal(tc.input); got != tc.want {
				t.Fatalf("NumberToOrdinal(%d) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}
