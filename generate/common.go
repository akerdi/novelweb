package generate

import (
	"crypto/md5"
	"encoding/hex"
)

func Ternary(expr bool, whenTrue, whenFalse interface{}) interface{} {
	if expr == true {
		return  whenTrue
	}
	return whenFalse
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}