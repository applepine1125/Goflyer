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

type CoinAddress struct {
	Type         string `json:"type"`
	CurrencyCode string `json:"currency_code"`
	Address      string `json:"address"`
}

func (a *API) GetCoinAddresses() (coinAddresses []CoinAddress, err error) {
	path := "/v1/me/getaddresses"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return coinAddresses, err
	}

	err = json.Unmarshal(byteSlice, &coinAddresses)

	return coinAddresses, err
}

type CoinDepositHistory struct {
	ID           int64   `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Address      string  `json:"address"`
	TxHash       string  `json:"tx_hash"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (a *API) GetCoinDepositHistories() (coinDepositHistories []CoinDepositHistory, err error) {
	path := "/v1/me/getcoinins"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return coinDepositHistories, err
	}

	err = json.Unmarshal(byteSlice, &coinDepositHistories)

	return coinDepositHistories, err
}

type CoinSendingHistory struct {
	ID            int64   `json:"id"`
	OrderID       string  `json:"order_id"`
	CurrencyCode  string  `json:"currency_code"`
	Amount        float64 `json:"amount"`
	Address       string  `json:"address"`
	TxHash        string  `json:"tx_hash"`
	Fee           float64 `json:"fee"`
	AdditionalFee float64 `json:"additional_fee"`
	Status        string  `json:"status"`
	EventDate     string  `json:"event_date"`
}

func (a *API) GetCoinSendingHistories() (coinSendingHistories []CoinSendingHistory, err error) {
	path := "/v1/me/getcoinouts"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return coinSendingHistories, err
	}

	err = json.Unmarshal(byteSlice, &coinSendingHistories)

	return coinSendingHistories, err
}

type BankAccount struct {
	ID            int64  `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountType   string `json:"account_type"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

func (a *API) GetBankAccounts() (bankAccounts []BankAccount, err error) {
	path := "/v1/me/getbankaccounts"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return bankAccounts, err
	}

	err = json.Unmarshal(byteSlice, &bankAccounts)

	return bankAccounts, err
}

type MoneyDepositHistory struct {
	ID           int64   `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (a *API) GetMoneyDepositHistories() (moneyDepositHistories []MoneyDepositHistory, err error) {
	path := "/v1/me/getdeposits"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return moneyDepositHistories, err
	}

	err = json.Unmarshal(byteSlice, &moneyDepositHistories)

	return moneyDepositHistories, err
}
