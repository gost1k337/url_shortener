package hasher

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

const SaltSize = 16

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)

	sha512Hasher := sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}
