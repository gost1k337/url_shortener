package hasher

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

const SaltSize = 16

func GenerateRandomSalt(saltSize int) ([]byte, error) {
	salt := make([]byte, saltSize)

	_, err := rand.Read(salt)
	if err != nil {
		return []byte{}, fmt.Errorf("read: %w", err)
	}

	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)

	sha512Hasher := sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}
