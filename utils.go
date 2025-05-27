package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
	"time"
)

func SecureRandomHex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)

}

func HmacHexDigest(hashFunc func() hash.Hash, secret, message string) (string, error) {
	h := hmac.New(hashFunc, []byte(secret))

	if _, err := h.Write([]byte(message)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func Map[T any, U any](list []T, mapFunc func(T) U) (mapped []U) {
	mapped = make([]U, len(list))
	for i, v := range list {
		mapped[i] = mapFunc(v)
	}

	return

}

func MapWithIndex[T any, U any](list []T, mapFunc func(T, int) U) (mapped []U) {
	mapped = make([]U, len(list))
	for i, v := range list {
		mapped[i] = mapFunc(v, i)
	}

	return

}

func StringToUint(i string) (parseInt uint) {
	parsedInt, err := strconv.ParseUint(i, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(parsedInt)
}

func Blank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func LeftPad(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}

	padCount := length - len(s)
	return strings.Repeat(string(pad), padCount) + s
}

func Pop[T any](list []T) (popped T) {
	if len(list) == 0 {
		panic("Cannot pop from an empty slice")
	}

	popped = list[len(list)-1]

	return
}

type Number interface{ int | float64 }

func Sum[T Number](list []T) (sum T) {
	for _, digit := range list {
		sum += digit
	}

	return
}

func ToDate(layout, date string) (parsedDate *time.Time) {
	if Blank(date) {
		return
	}

	t, err := time.Parse(layout, date)
	if err != nil {
		return
	}

	return &t

}
