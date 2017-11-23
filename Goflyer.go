package Goflyer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type API struct {
	secret string
	key    string
	client *http.Client
}

func NewAPI(secret string, key string) (a *API) {
	a = new(API)
	a.secret = secret
	a.key = key
	a.client = new(http.Client)
	return a
}

func (a *API) Request(url string, endPoint string, method string, params string) ([]byte, error) {
	req, _ := http.NewRequest(method, url+endPoint+params, nil)
	resp, err := a.client.Do(req)

	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returns status %s", resp.Status)
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return byteArray, err
	}

	return byteArray, nil
}
