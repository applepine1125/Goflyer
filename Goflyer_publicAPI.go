package Goflyer

import (
	"encoding/json"
)

type Market struct {
	ProductCode string `json:"product_code"`
	Alias       string `json:"alias"`
}

func (a *API) GetMarkets() ([]Market, error) {
	path := "/v1/markets"
	var market []Market

	byteArray, err := a.PublicAPIRequest(path, "GET", "")
	if err != nil {
		return market, err
	}

	err = json.Unmarshal(byteArray, &market)

	return market, err
}

type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

type Order struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

func (a *API) GetBoard(ProductCode string) (Board, error) {
	path := "/v1/board"
	params := ""
	if ProductCode != "" {
		params = "?product_code=" + ProductCode
	}
	var board Board

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return board, err
	}

	err = json.Unmarshal(byteArray, &board)

	return board, err
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	TimeStamp       string  `json:"timestamp"`
	TickID          int64   `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (a *API) GetTicker(ProductCode string) (Ticker, error) {
	path := "/v1/ticker"
	params := ""
	if ProductCode != "" {
		params = "?product_code=" + ProductCode
	}
	var ticker Ticker

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return ticker, err
	}
	err = json.Unmarshal(byteArray, &ticker)

	return ticker, err
}

type Execution struct {
	ID                         int64   `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_d"`
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

func (a *API) GetExecutions(ProductCode string) ([]Execution, error) {
	path := "/v1/executions"
	params := ""
	if ProductCode != "" {
		params = "?product_code=" + ProductCode
	}
	var executions []Execution

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return executions, err
	}
	err = json.Unmarshal(byteArray, &executions)

	return executions, err
}

type BoardState struct {
	Health string           `json:"health"`
	State  string           `json:"state"`
	Data   specialQuotation `json:"data"`
}

type specialQuotation struct {
	SpecialQuotation float64 `json:"special_quotation"`
}

func (a *API) GetBoardState(ProductCode string) (BoardState, error) {
	path := "/v1/getboardstate"
	params := ""
	if ProductCode != "" {
		params = "?product_code=" + ProductCode
	}
	var boardState BoardState

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return boardState, err
	}
	err = json.Unmarshal(byteArray, &boardState)

	return boardState, err
}

type MarketHealth struct {
	Status string `json:"status"`
}

func (a *API) GetMarketHealth(ProductCode string) (MarketHealth, error) {
	path := "/v1/gethealth"
	params := ""
	if ProductCode != "" {
		params = "?product_code=" + ProductCode
	}

	var marketHealth MarketHealth

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return marketHealth, err
	}
	err = json.Unmarshal(byteArray, &marketHealth)

	return marketHealth, err
}

type ChatMessage struct {
	NickName string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

func (a *API) GetChatMessages(fromDate string) ([]ChatMessage, error) {
	path := "/v1/getchats"
	params := ""
	if fromDate != "" {
		params = "?from_date=" + fromDate
	}
	var chatMessages []ChatMessage

	byteArray, err := a.PublicAPIRequest(path, "GET", params)
	if err != nil {
		return chatMessages, err
	}
	err = json.Unmarshal(byteArray, &chatMessages)

	return chatMessages, err
}