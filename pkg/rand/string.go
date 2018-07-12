package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateString generate a random string using `rand.Read()`
func GenerateString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
