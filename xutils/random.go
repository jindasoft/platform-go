package xutils

import (
	"crypto/rand"
)

func Random(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)

	for i := range length {
		randomIndex := make([]byte, 1)
		rand.Read(randomIndex)
		result[i] = chars[int(randomIndex[0])%len(chars)]
	}

	return string(result)
}

func RandomLower(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)

	for i := range length {
		randomIndex := make([]byte, 1)
		rand.Read(randomIndex)
		result[i] = chars[int(randomIndex[0])%len(chars)]
	}

	return string(result)
}

func RandomUpper(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := range length {
		randomIndex := make([]byte, 1)
		rand.Read(randomIndex)
		result[i] = chars[int(randomIndex[0])%len(chars)]
	}

	return string(result)
}

func RandomUNumber(length int) string {
	chars := "0123456789"
	result := make([]byte, length)

	for i := range length {
		randomIndex := make([]byte, 1)
		rand.Read(randomIndex)
		result[i] = chars[int(randomIndex[0])%len(chars)]
	}

	return string(result)
}
