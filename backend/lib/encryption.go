package lib

import "crypto/sha512"

func HashString(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	return sliceHashToByteArray(sha512.Sum512([]byte(s)))
}

func CompareHashAndString(h []byte, s string) bool {
	return string(h) == string(HashString(s))
}

func sliceHashToByteArray(s [64]byte) []byte {
	return s[:]
}
