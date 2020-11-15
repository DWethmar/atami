package message

import (
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

/**
Stuff to read:
https://blog.golang.org/strings
*/

// ParseResult result
type ParseResult struct {
	NormalizedText string
	WeightedLength int
}

// ParseMessage return info about the text length
func ParseMessage(txt string) ParseResult {
	normalizedText := norm.NFC.String(txt)
	// w := 0
	// for i := 0; i < len(normalizedText); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(txt[i:])
	// 	// _, width := utf8.DecodeRuneInString(txt[i:])
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	// 	w = width
	// }
	return ParseResult{
		NormalizedText: normalizedText,
		WeightedLength: utf8.RuneCountInString(normalizedText),
	}
}
