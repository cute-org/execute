package auth

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"io"
)

const (
	argonTime    = 1         // Iterations
	argonMemory  = 64 * 1024 // 64 MB in KiB
	argonThreads = 4         // Parallelism
	argonKeyLen  = 32        // Hash length
	saltLen      = 16        // Salt length
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltLen)
	_, err := io.ReadFull(rand.Reader, salt)
	return salt, err
}

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
	return base64.StdEncoding.EncodeToString(hash)
}

func EncodeSalt(salt []byte) string {
	return base64.StdEncoding.EncodeToString(salt)
}

func DecodeSalt(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

func CompareHashes(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	result := 0
	for i := range a {
		result |= int(a[i]) ^ int(b[i])
	}
	return result == 0
}
