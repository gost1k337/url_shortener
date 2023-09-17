package hasher

import (
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	grow     = 10
)

func NewShortURL(number uint64) string {
	var encodedBuilder strings.Builder

	length := len(alphabet)

	encodedBuilder.Grow(grow)

	for ; number > 0; number /= uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}
