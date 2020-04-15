package generate

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

func Ternary(expr bool, whenTrue, whenFalse interface{}) interface{} {
	if expr == true {
		return whenTrue
	}
	return whenFalse
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func UrlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	return baseURL.ResolveReference(uri).String()
}
