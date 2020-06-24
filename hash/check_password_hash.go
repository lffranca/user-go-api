package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash CheckPasswordHash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
