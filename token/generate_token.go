package token

import (
	"crypto/rand"
	"fmt"
)

// GenerateToken GenerateToken
func GenerateToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
