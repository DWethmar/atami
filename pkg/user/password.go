package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password
func HashPassword(plainPwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(plainPwd, bcrypt.MaxCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compairs a hashed password with a plain text password.
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
