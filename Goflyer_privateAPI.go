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

func (a *API) GetBalance() (balances []Balance, err error) {
	path := "/v1/me/getbalance"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return balances, err
	}

	err = json.Unmarshal(byteSlice, &balances)

	return balances, err
}

type Collateral struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `jdon:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

func (a *API) GetCollateral() (collateral Collateral, err error) {
	path := "/v1/me/getcollateral"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return collateral, err
	}

	err = json.Unmarshal(byteSlice, &collateral)

	return collateral, err
}

type CollateralAccount struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

func (a *API) GetCollateralAccounts() (collateralAccounts []CollateralAccount, err error) {
	path := "/v1/me/getcollateralaccounts"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return collateralAccounts, err
	}

	err = json.Unmarshal(byteSlice, &collateralAccounts)

	return collateralAccounts, err
}

type Address struct {
	Type         string `json:"type"`
	CurrencyCode string `json:"currency_code"`
	Address      string `json:"address"`
}

func (a *API) GetAddresses() (addresses []Address, err error) {
	path := "/v1/me/getaddresses"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return addresses, err
	}

	err = json.Unmarshal(byteSlice, &addresses)

	return addresses, err
}

type DepositHistory struct {
	ID           int64   `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Address      string  `json:"address"`
	TxHash       string  `json:"tx_hash"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (a *API) GetDepositHistories() (depositHistories []DepositHistory, err error) {
	path := "/v1/me/getcoinins"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return depositHistories, err
	}

	err = json.Unmarshal(byteSlice, &depositHistories)

	return depositHistories, err
}
