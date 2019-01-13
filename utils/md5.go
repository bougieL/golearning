package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 hash
func Md5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
