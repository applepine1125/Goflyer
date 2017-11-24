package Goflyer

import (
	"encoding/json"
	"strings"
)

func (a *API) GetPermissions() ([]string, error) {
	path := "/v1/me/getpermissions"

	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")
	return byteSlice2strSlice(byteSlice), err
}

func byteSlice2strSlice(bs []byte) []string {
	ss := strings.SplitAfter(string(bs), "\",\"")

	for i, _ := range ss {
		ss[i] = strings.TrimSuffix(ss[i], "\",\"")
		ss[i] = strings.TrimPrefix(ss[i], "[\"")
		ss[i] = strings.TrimSuffix(ss[i], "\"]")
	}

	return ss
}

type Balance struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}

func (a *API) GetBalance() ([]Balance, error) {
	path := "/v1/me/getbalance"

	var balances []Balance
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return balances, err
	}

	err = json.Unmarshal(byteSlice, &balances)

	return balances, err
}
