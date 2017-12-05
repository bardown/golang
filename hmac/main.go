package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func computeHmacSha1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func sendPayLoad(resource, password string) {
	apiURL := "https://jsonplaceholder.typicode.com"
	resource = "/users/"
	data := url.Values{}
	data.Set("name", "foo")
	data.Add("surname", "bar")

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("Basic: %s", password))
	r.Header.Add("Content-:Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}


func main() {
	body := []string{"time.Now(
	sendPayLoad("/users/")
}
