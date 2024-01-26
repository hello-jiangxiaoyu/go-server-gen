package utils

import (
	"crypto/rand"
	"github.com/google/uuid"
	"strings"
)

const base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const hexChars = "0123456789abcdef"
const decChars = "0123456789"

func Rand64(n int) string  { return RandString(n, base64Chars) }
func Rand62(n int) string  { return RandString(n, base62Chars) }
func RandDec(n int) string { return RandString(n, decChars) }
func RandHex(n int) string { return RandString(n, hexChars) }

// RandString Generate random string
func RandString(n int, letters string) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	if len(letters) == 0 {
		letters = base62Chars
	}
	for i := 0; i < len(b); i++ {
		b[i] = letters[b[i]%byte(len(letters))]
	}

	return string(b)
}

func GetNoLineUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
