package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex")

func Sha256Hasher(userString string) string {
	hash_:=sha256.Sum256([]byte(userString))
	encodedString:=hex.EncodeToString(hash_[:])
	return encodedString
}

func Md5Hasher(userString string) string{
	hash_:=md5.Sum([]byte(userString))
	encodedString:=hex.EncodeToString(hash_[:])
	return encodedString
}