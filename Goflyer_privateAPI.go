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

type MoneyWithdrawResp struct {
	MessageID string `json:"message_id"`
}

type MoneyWithdrawBody struct {
	CurrencyCode  string `json:"currency_code"`
	BankAccountID int64  `json:"bank_account_id"`
	Amount        int64  `json:"amount"`
	Code          int64  `json:"code"`
}

func (a *API) MoneyWithdraw(moneyWithdrawBody MoneyWithdrawBody) (moneyWithdrawResp MoneyWithdrawResp, err error) {
	path := "/v1/me/withdraw"

	jsonByte, err := json.Marshal(moneyWithdrawBody)
	if err != nil {
		return moneyWithdrawResp, err
	}

	byteSlice, err := a.PrivateAPIRequest(path, "POST", string(jsonByte))
	if err != nil {
		return moneyWithdrawResp, err
	}

	err = json.Unmarshal(byteSlice, &moneyWithdrawResp)
	return moneyWithdrawResp, err
}

type MoneyWithdrawHistory struct {
	ID           int64  `json:"id"`
	OrderID      string `json:"order_id"`
	CurrencyCode string `json:"currency_code"`
	Amount       int64  `json:"amount"`
	Status       string `json:"status"`
	EventDate    string `json:"event_date"`
}

func (a *API) GetMoneyWithdrawHistories() (moneyWithdrawHistories []MoneyWithdrawHistory, err error) {
	path := "/v1/me/getwithdrawals"
	byteSlice, err := a.PrivateAPIRequest(path, "GET", "")

	if err != nil {
		return moneyWithdrawHistories, err
	}

	err = json.Unmarshal(byteSlice, &moneyWithdrawHistories)
	return moneyWithdrawHistories, err
}

type SendChildOrderBody struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          float64 `json:"price"`
	Size           float64 `json:"size"`
	MinutetoExpire int64   `json:"minute_to_expire"`
	TimeinForce    string  `json:"time_in_force"`
}

type SendChildOrderResp struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func (a *API) SendChildOrder(sendChildOrderBody SendChildOrderBody) (sendChildOrderResp SendChildOrderResp, err error) {
	path := "/v1/me/sendchildorder"

	jsonByte, err := json.Marshal(sendChildOrderBody)
	if err != nil {
		return sendChildOrderResp, err
	}

	byteSlice, err := a.PrivateAPIRequest(path, "POST", string(jsonByte))
	if err != nil {
		return sendChildOrderResp, err
	}

	err = json.Unmarshal(byteSlice, &sendChildOrderResp)
	return sendChildOrderResp, err
}

type CancelChildOrderBody struct {
	ProductCode  string `json:"product_code"`
	ChildOrderID string `json:"child_order_id"`
}

func (a *API) CancelChildOrder(cancelChildOrderBody CancelChildOrderBody) (err error) {
	path := "/v1/me/cancelchildorder"

	jsonByte, err := json.Marshal(cancelChildOrderBody)
	if err != nil {
		return err
	}

	_, err = a.PrivateAPIRequest(path, "POST", string(jsonByte))
	return err
}

type SendParentOrderBody struct {
	OrderMethod    string                     `json:"order_method"`
	MinutetoExpire int64                      `json:"minute_to_expire"`
	TimeinForce    string                     `json:"time_in_force"`
	Params         []SendParentOrderBodyParam `json:"parameters"`
}

type SendParentOrderBodyParam struct {
	ProductCode   string  `json:"product_code"`
	ConditionType string  `json:"condition_type"`
	Side          string  `json:"side"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
}

type SendParentOrderResp struct {
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id"`
}

func (a *API) SendParentOrder(sendParentOrderBody SendParentOrderBody) (sendParentOrderResp SendParentOrderResp, err error) {
	path := "/v1/me/sendparentorder"

	jsonByte, err := json.Marshal(sendParentOrderBody)
	if err != nil {
		return sendParentOrderResp, err
	}

	byteSlice, err := a.PrivateAPIRequest(path, "POST", string(jsonByte))
	if err != nil {
		return sendParentOrderResp, err
	}

	err = json.Unmarshal(byteSlice, &sendParentOrderResp)
	return sendParentOrderResp, err
}

type CancelParentOrderBody struct {
	ProductCode   string `json:"product_code"`
	ParentOrderID string `json:"parent_order_id"`
}

func (a *API) CancelParentOrder(cancelParentOrderBody CancelParentOrderBody) (err error) {
	path := "/v1/me/cancelparentorder"

	jsonByte, err := json.Marshal(cancelParentOrderBody)
	if err != nil {
		return err
	}

	_, err = a.PrivateAPIRequest(path, "POST", string(jsonByte))
	return err
}

type CancelAllChildOrdersBody struct {
	ProductCode string `json:"product_code"`
}

func (a *API) CancelAllChildOrders(cancelAllChildOrdersBody CancelAllChildOrdersBody) (err error) {
	path := "/v1/me/cancelallchildorders"

	jsonByte, err := json.Marshal(cancelAllChildOrdersBody)
	if err != nil {
		return err
	}

	_, err = a.PrivateAPIRequest(path, "POST", string(jsonByte))
	return err
}

type ChildOrder struct {
	ID                     int64   `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	ProductCode            string  `json:"product_code"`
	Side                   string  `json:"side"`
	ChildOrderType         string  `json:"child_order_type"`
	Price                  float64 `json:"price"`
	AveragePrice           float64 `json:"average_price"`
	Size                   float64 `json:"size"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommision         float64 `json:"total_commission"`
}

func (a *API) GetChildOrders(query string) (childOrders []ChildOrder, err error) {
	path := "/v1/me/getchildorders"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return childOrders, err
	}

	err = json.Unmarshal(byteSlice, &childOrders)
	return childOrders, err
}

type ParentOrder struct {
	ID                      int64   `json:"id"`
	ParentOrderID           string  `json:"parent_order_id"`
	ProductCode             string  `json:"product_code"`
	Side                    string  `json:"side"`
	ParentOrderType         string  `json:"parent_order_type"`
	Price                   float64 `json:"price"`
	AveragePrice            float64 `json:"average_price"`
	Size                    float64 `json:"size"`
	ParentOrderState        string  `json:"parent_order_state"`
	ExpireDate              string  `json:"expire_date"`
	ParentOrderDate         string  `json:"parent_order_date"`
	ParentOrderAcceptanceID string  `json:"parent_order_acceptance_id"`
	OutstandingSize         float64 `json:"outstanding_size"`
	CancelSize              float64 `json:"cancel_size"`
	ExecutedSize            float64 `json:"executed_size"`
	TotalCommision          float64 `json:"total_commission"`
}

func (a *API) GetParentOrders(query string) (parentOrders []ParentOrder, err error) {
	path := "/v1/me/getparentorders"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return parentOrders, err
	}

	err = json.Unmarshal(byteSlice, &parentOrders)
	return parentOrders, err
}

type ParentOrderDetail struct {
	ID                      int64                    `json:"id"`
	ParentOrderID           string                   `json:"parent_order_id"`
	OrderMethod             string                   `json:"order_method"`
	MinutetoExpire          int64                    `json:"minute_to_expire"`
	Parameters              []ParentOrderDetailParam `json:"parameters"`
	ParentOrderAcceptanceID string                   `json:"parent_order_acceptance_id"`
}

type ParentOrderDetailParam struct {
	ProductCode   string  `json:"product_code"`
	ConditionType string  `json:"condition_type"`
	Side          string  `json:"side"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	TriggerPrice  float64 `json:"trigger_price"`
	Offset        float64 `json:"offset"`
}

func (a *API) GetParentOrderDetail(query string) (parentOrderDetail ParentOrderDetail, err error) {
	path := "/v1/me/getparentorder"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return parentOrderDetail, err
	}

	err = json.Unmarshal(byteSlice, &parentOrderDetail)
	return parentOrderDetail, err
}

type Execution struct {
	ID                     int64   `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	Commission             float64 `json:"commission"`
	ExecDate               string  `json:"exec_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
}

func (a *API) GetExecutions(query string) (executions []Execution, err error) {
	path := "/v1/me/getexecuions"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return executions, err
	}

	err = json.Unmarshal(byteSlice, &executions)
	return executions, err
}

type Position struct {
	ProductCode         string  `json:"product_code"`
	Side                string  `json:"side"`
	Price               float64 `json:"price"`
	Size                float64 `json:"size"`
	Commission          float64 `json:"commission"`
	SwapPointAccumulate float64 `json:"swap_point_accumulate"`
	RequireCollateral   float64 `json:"require_collateral"`
	OpenDate            string  `json:"open_date"`
	Leverage            float64 `json:"leverage"`
	Pnl                 float64 `json:"pnl"`
}

func (a *API) GetPositions(productCode string) (positions []Position, er error) {
	path := "/v1/me/getpositions"
	params := ""
	if productCode != "" {
		params = "?product_code=" + productCode
	}

	byteSlice, err := a.PrivateAPIRequest(path+params, "GET", "")
	if err != nil {
		return positions, err
	}

	err = json.Unmarshal(byteSlice, &positions)
	return positions, err
}

type CollateralHistory struct {
	ID           int64   `json:"id"`
	CurrencyCode string  `json:"currency_code"`
	Change       float64 `json:"change"`
	Amount       float64 `json:"amount"`
	ReasonCode   string  `json:"reason_code"`
	Date         string  `json:"date"`
}

func (a *API) GetCollateralHistories(query string) (collateralHistories []CollateralHistory, err error) {
	path := "/v1/me/getcollateralhistory"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return collateralHistories, err
	}

	err = json.Unmarshal(byteSlice, &collateralHistories)
	return collateralHistories, err
}

type TradingCommission struct {
	CommissionRate float64 `json:"commission_rate"`
}

func (a *API) GetCommissionRate(query string) (tradingCommission []TradingCommission, err error) {
	path := "/v1/me/gettradingcommission"

	byteSlice, err := a.PrivateAPIRequest(path+query, "GET", "")
	if err != nil {
		return tradingCommission, err
	}

	err = json.Unmarshal(byteSlice, &tradingCommission)
	return tradingCommission, err
}
