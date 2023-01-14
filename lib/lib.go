package lib

import (
	"crypto/sha512"
	"fmt"
)

func MakeHash(p string) string {
	pbyte := []byte(p)
	pHash := sha512.Sum512(pbyte)
	return fmt.Sprintf("%x", pHash)
}
