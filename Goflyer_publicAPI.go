package Goflyer

import (
	"encoding/json"
)

type Market struct {
	ProductCode string `json:"product_code"`
	Alias       string `json:"alias"`
}

func (a *API) GetMarkets() (markets []Market, err error) {
	path := "/v1/markets"

	byteSlice, err := a.PublicAPIRequest(path, "GET", "")
	if err != nil {
		return markets, err
	}

	err = json.Unmarshal(byteSlice, &markets)
	return markets, err
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

func (a *API) GetBoard(query string) (board Board, err error) {
	path := "/v1/board"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return board, err
	}

	err = json.Unmarshal(byteSlice, &board)
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

func (a *API) GetTicker(query string) (ticker Ticker, err error) {
	path := "/v1/ticker"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return ticker, err
	}

	err = json.Unmarshal(byteSlice, &ticker)
	return ticker, err
}

type MarketExecution struct {
	ID                         int64   `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_d"`
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

func (a *API) GetMarketExecutions(query string) (marketExecutions []MarketExecution, err error) {
	path := "/v1/executions"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return marketExecutions, err
	}

	err = json.Unmarshal(byteSlice, &marketExecutions)
	return marketExecutions, err
}

type BoardState struct {
	Health string           `json:"health"`
	State  string           `json:"state"`
	Data   specialQuotation `json:"data"`
}

type specialQuotation struct {
	SpecialQuotation float64 `json:"special_quotation"`
}

func (a *API) GetBoardState(query string) (boardState BoardState, err error) {
	path := "/v1/getboardstate"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return boardState, err
	}

	err = json.Unmarshal(byteSlice, &boardState)
	return boardState, err
}

type MarketHealth struct {
	Status string `json:"status"`
}

func (a *API) GetMarketHealth(query string) (marketHealth MarketHealth, err error) {
	path := "/v1/gethealth"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return marketHealth, err
	}

	err = json.Unmarshal(byteSlice, &marketHealth)
	return marketHealth, err
}

type ChatMessage struct {
	NickName string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

func (a *API) GetChatMessages(query string) (chatMessages []ChatMessage, err error) {
	path := "/v1/getchats"
	byteSlice, err := a.PublicAPIRequest(path, "GET", query)
	if err != nil {
		return chatMessages, err
	}

	err = json.Unmarshal(byteSlice, &chatMessages)
	return chatMessages, err
}
