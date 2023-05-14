package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
	{7, "7"},
	{8, "8"},
	{9, "Fizz"},
	{10, "Buzz"},
	{11, "11"},
	{12, "Fizz"},
	{13, "13"},
	{14, "14"},
	{15, "FizzBuzz"},
	{16, "16"},
	{17, "17"},
	{18, "Fizz"},
	{19, "19"},
	{20, "Buzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range testCases {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		testName := fmt.Sprintf("should return %s when input is %d", tt.expected, tt.input)

		t.Run(testName, func(t *testing.T) {
			// t.Parallel()
			// When
			result := FizzBuzz(tt.input)
			// Then
			assert.Equal(t, tt.expected, result)
		})
	}
}
