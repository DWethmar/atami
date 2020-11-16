package msgtxt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHashtags(t *testing.T) {
	tests := []struct {
		Text     string
		Hashtags []string
	}{
		{
			"Test test #hashtag1 test test #hashtag2",
			[]string{"hashtag1", "hashtag2"},
		},
		{
			"#hashtag1 test test#hashtag2",
			[]string{"hashtag1"},
		},
		{
			"#💩 test test#hashtag2",
			[]string{},
		},
		{
			"#test#",
			[]string{"test"},
		},
		{
			"#test#test#hasttags",
			[]string{"test"},
		},
	}

	for _, test := range tests {
		r := FindHashtags(test.Text)

		fmt.Printf("expect: %d received: %d TXT: %s \n", len(test.Hashtags), len(r), test.Text)

		if assert.Equal(t, len(test.Hashtags), len(r), test.Text) {
			for i, hashtag := range test.Hashtags {
				assert.Equal(t, test.Hashtags[i], hashtag, test.Text)
			}
		}
	}
}

// ParseMessage return info about the text length
// go test ./pkg/message/ -v
func TestParse(t *testing.T) {
	tests := []struct {
		Text   string
		Length int
	}{
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
