package internal

import (
	"golang.org/x/crypto/argon2"
)

// generatePasswordHash computes the Argon2id hash using a given password and salt
func generatePasswordHash(password string, salt []byte) []byte {
	// Parameters: 1 iteration, 64 MB memory, 4 threads, 32-byte key length
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

// compareHashes performs a constant-time comparison between two Base64-encoded hash strings
func compareHashes(hashA, hashB string) bool {
	if len(hashA) != len(hashB) {
		return false
	}
	result := 0
	for i, charA := range hashA {
		result |= int(charA) ^ int(hashB[i])
	}
	return result == 0
}
