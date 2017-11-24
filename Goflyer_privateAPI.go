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

type Collateral struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `jdon:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

func (a *API) GetCollateral() (Collateral, error) {
	path := "/v1/me/getcollateral"

	var collateral Collateral
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

func (a *API) GetCollateralAccounts() ([]CollateralAccount, error) {
	path := "/v1/me/getcollateralaccounts"

	var collateralAccounts []CollateralAccount
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return collateralAccounts, err
	}

	err = json.Unmarshal(byteSlice, &collateralAccounts)

	return collateralAccounts, err
}
