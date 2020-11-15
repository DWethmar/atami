package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type compareLength struct {
	Text   string
	Length int
}

// ParseMessage return info about the text length
func TestParseMessage(t *testing.T) {
	tests := []compareLength{
		{
			"👦",
			1,
		},
		{
			"Amélie",
			6,
		},
		{
			"\u0041\u006d\u00e9\u006c\u0069\u0065", // Amélie
			6,
		},
		{
			"\u0041\u006d\u0065\u0301\u006c\u0069\u0065", // Amélie
			6,
		},
		{
			"https://www.youtube.com/watch?v=dQw4w9WgXcQ thats a cool link",
			61,
		},
		{
			"おはようございます。",
			10,
		},
	}

	for _, test := range tests {
		r := ParseMessage(test.Text)
		fmt.Printf("expect: %d received: %d TXT: %s \n", test.Length, r.WeightedLength, test.Text)
		assert.Equal(t, test.Length, r.WeightedLength, fmt.Sprintf("B length: %d TXT: %s", r.WeightedLength, test.Text))
	}
}
