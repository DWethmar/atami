package msgtxt

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
// go test ./pkg/message/ -v
func TestParse(t *testing.T) {
	tests := []compareLength{
		{
			"👨‍👩‍👧‍👦",
			7,
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
		{
			"\u00F1",
			1,
		},
		{
			"'\u006E\u0303",
			2,
		},
	}

	for _, test := range tests {
		r := Parse(test.Text)
		fmt.Printf("expect: %d received: %d TXT: %s \n", test.Length, r.NormalizedLength, test.Text)
		assert.Equal(t, test.Length, r.NormalizedLength, fmt.Sprintf("B length: %d TXT: %s", r.NormalizedLength, test.Text))
	}
}
