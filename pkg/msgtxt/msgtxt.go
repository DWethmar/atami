package msgtxt

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

/**
Stuff to read:
https://blog.golang.org/strings
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/normalize
https://developer.twitter.com/en/docs/counting-characters#:~:text=Definition%20of%20a%20Character,280%20characters%20or%20Unicode%20glyphs.
https://www.objc.io/blog/2017/12/19/decomposing-emoji/
https://developer.twitter.com/en/docs/tco
*/

// ParseResult result
type ParseResult struct {
	NormalizedText   string
	NormalizedLength int
	Hashtags         []string
}

// FindHashtags finds hastgas in text
func FindHashtags(txt string) []string {
	hashtags := make([]string, 0)
	var re = regexp.MustCompile(`\B\#\w\w+\b`)

	for i, match := range re.FindAllString(txt, -1) {
		fmt.Println(match, "found at index", i)
		hashtags = append(hashtags, match)
	}

	return hashtags
}

// Parse return info about the text length
func Parse(txt string) ParseResult {
	normalizedText := norm.NFC.String(txt)

	return ParseResult{
		NormalizedText:   normalizedText,
		NormalizedLength: utf8.RuneCountInString(normalizedText),
	}
}
