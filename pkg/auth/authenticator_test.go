package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var passwords = []string{
	"",
	"nice",
	"verynice",
	"NjY5ZjVlNTg0MThkNTZjODNjNjQ0NWEx",
	"a",
	"1",
	"this is a password with spaces in it",
	"ECabx;^RLk&-hW5q*J4*Q",
	"7+h<+%c{<XBT^%$9mnm_5",
	"+X&Ah9kwhW[bQ:k]p^q4Y",
	"9pRuGP[x9:Z\\E`pmY=7VS",
	"t/;_BxR!W%-.Uk\\L#)(=[CdEg.eEZ%Ha}-\"vc[m'.]{!E(Kkx;s%Vvt&dka{_F@56@-3v7p#?cv3[h-E]9?k&a^*,!!-Z{HnBtqh2:DhhSt9;vr5eAJ,",
}

// TestAuthenticate tests the Delete function.
func TestDefaultComparePassword(t *testing.T) {
	for _, password := range passwords {
		hashedPassword := HashPassword([]byte(password))
		assert.True(t, ComparePasswords(hashedPassword, []byte(password)), fmt.Sprintf("Password %s with hash %s are not equal", password, hashedPassword))
	}
}

func TestInvalidDefaultComparePassword(t *testing.T) {
	for _, password := range passwords {
		hashedPassword := HashPassword([]byte(password))
		assert.False(t, ComparePasswords(hashedPassword, []byte("oof")), fmt.Sprintf("Password %s with hash %s are equal", password, hashedPassword))
	}
}
