package util

import (
	"crypto/rand"
	"math/big"
)

const ALPHABET_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
const ALPHABET_NUMERIC = "0123456789"
const ALPHABET_SPECIAL = "!@#$%^&*()_+-=[]{};':\",./<>?"
const ALPHABET_ALL = ALPHABET_UPPER + ALPHABET_LOWER + ALPHABET_NUMERIC + ALPHABET_SPECIAL
const ALPHABET_URL_SAFE = ALPHABET_UPPER + ALPHABET_LOWER + ALPHABET_NUMERIC + "-_"

func CryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func RandomString(length int, alphabets []string) string {
	token := ""
	codeAlphabet := ""
	for _, alphabet := range alphabets {
		codeAlphabet += alphabet
	}
	for i := 0; i < length; i++ {
		token += string(codeAlphabet[CryptoRandSecure(int64(len(codeAlphabet)))])
	}
	return token
}
