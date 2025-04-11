package auth

import (
	"golang.org/x/crypto/argon2"
)

const (
	argonTime    = 1         // Number of iterations
	argonMemory  = 64 * 1024 // 64 MB of memory in KiB
	argonThreads = 4         // Number of threads (parallelism)
	argonKeyLen  = 32        // Length of the generated hash in bytes
)

// generatePasswordHash computes the Argon2id hash using a given password and salt
func generatePasswordHash(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
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
