package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"
)

func computeHmacSha1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {

	now := time.Now()
	method := "GET"
	url := "https://supersecureapi"
	extra := "/admin/v1/users"

	conan := fmt.Sprintf("%s\n%s\n%s\n%s\n", now, method, url, extra)
	computeHmacSha1(conan, "lkjj1l2kj3ajaf")

	fmt.Println(conan)
	fmt.Println(computeHmacSha1)

}
