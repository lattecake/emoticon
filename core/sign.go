package core

import (
	"crypto/md5"
	"encoding/hex"
)

func Sign(name string) string {
	h := md5.New()
	h.Write([]byte(name))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

