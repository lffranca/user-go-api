package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash GenerateHash
func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
