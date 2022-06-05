package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMD5Hash(v string) string {
	hash := md5.Sum([]byte(v))

	return hex.EncodeToString(hash[:])
}
