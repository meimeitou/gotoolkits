package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(o string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(o))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
