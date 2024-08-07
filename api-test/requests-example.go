package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	balance()
	deposit()
	checkWallet()
	monthHistory()
}

func balance() {
	url := "http://localhost:8080/wallet/balance"
	req, _ := http.NewRequest("GET", url, strings.NewReader(""))

	reqBody, _ := io.ReadAll(req.Body)
	hash := hmac.New(sha1.New, []byte("mySecretKey"))
	hash.Write(reqBody)
	actualMac := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("X-Digest", actualMac)
	req.Header.Add("X-UserId", "1")

	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func deposit() {
	url := "http://localhost:8080/wallet/deposit"
	req, _ := http.NewRequest("POST", url, strings.NewReader(`{
		"amount": 199
	}`))

	reqBody, _ := io.ReadAll(req.Body)
	hash := hmac.New(sha1.New, []byte("mySecretKey"))
	hash.Write(reqBody)
	actualMac := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("X-Digest", actualMac)
	req.Header.Add("X-UserId", "1")

	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func checkWallet() {
	url := "http://localhost:8080/wallet/check-account"
	req, _ := http.NewRequest("GET", url, strings.NewReader(``))

	reqBody, _ := io.ReadAll(req.Body)
	hash := hmac.New(sha1.New, []byte("mySecretKey"))
	hash.Write(reqBody)
	actualMac := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("X-Digest", actualMac)
	req.Header.Add("X-UserId", "1")

	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func monthHistory() {
	url := "http://localhost:8080/wallet/month-history"
	req, _ := http.NewRequest("GET", url, strings.NewReader(``))

	reqBody, _ := io.ReadAll(req.Body)
	hash := hmac.New(sha1.New, []byte("mySecretKey"))
	hash.Write(reqBody)
	actualMac := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("X-Digest", actualMac)
	req.Header.Add("X-UserId", "1")

	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
