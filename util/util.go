package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5WithSalt(password, salt string) string {
	h := md5.New()

	h.Reset()
	bytes := []byte(salt)
	h.Write(bytes)
	saltArr := h.Sum(nil)

	h.Reset()
	h.Write(saltArr)

	h.Write([]byte(password))
	arr := h.Sum(nil)

	for i := 0; i < 1023; i++ {
		h.Reset()
		h.Write(arr)
		arr = h.Sum(nil)
	}
	return hex.EncodeToString(arr)
}
