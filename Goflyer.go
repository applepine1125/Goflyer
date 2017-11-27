package Goflyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type API struct {
	url    string
	key    string
	secret string
	client *http.Client
}

func NewAPI(key string, secret string) (a *API) {
	a = new(API)
	a.url = "https://api.bitflyer.jp"
	a.key = key
	a.secret = secret
	a.client = new(http.Client)
	return a
}

func (a *API) Request(req *http.Request) ([]byte, error) {
	resp, err := a.client.Do(req)

	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returns status %s", resp.Status)
	}

	defer resp.Body.Close()

	byteSlice, err := ioutil.ReadAll(resp.Body)

	return byteSlice, err
}

func (a *API) PublicAPIRequest(path string, method string, query string) ([]byte, error) {
	req, _ := http.NewRequest(method, a.url+path+query, nil)
	byteSlice, err := a.Request(req)

	return byteSlice, err
}

func (a *API) PrivateAPIRequest(path string, method string, query string, body string) ([]byte, error) {
	accessTimeStamp := string(time.Now().Unix())
	accessSign := accessTimeStamp + method + path + query + body
	hmac := hmac.New(sha256.New, []byte(a.secret))
	hmac.Write([]byte(accessSign))
	sha256AccessSign := hex.EncodeToString(hmac.Sum(nil))

	req, _ := http.NewRequest(method, a.url+path, nil)
	req.Header.Set("ACCESS-KEY", a.key)
	req.Header.Set("ACCESS-TIMESTAMP", accessTimeStamp)
	req.Header.Set("ACCESS-SIGN", sha256AccessSign)

	byteSlice, err := a.Request(req)

	return byteSlice, err
}
