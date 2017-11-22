package Goflyer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// API struct
type API struct {
	secret string
	key    string
	client *http.Client
}

//NewAPI returns API struct
func NewAPI(secret string, key string) (a *API) {
	a = new(API)
	a.secret = secret
	a.key = key
	a.client = new(http.Client)
	return a
}

// Request method in API
func (a *API) Request(url string, endPoint string, method string, params string) []byte {
	req, _ := http.NewRequest(method, url+endPoint+params, nil)

	resp, err := a.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return byteArray
}

type ohlcResult struct {
	Terms [][]float64 `json:"1800"`
	// CloseTime  *string
	// OpenPrice  *float64
	// HighPrice  *float64
	// LowPrice   *float64
	// ClosePrice *float64
	// Volume     *float64
}
type ohlcAllowance struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

// OHLC struct from cryptwatch
type OHLC struct {
	Result    ohlcResult
	Allowance ohlcAllowance
}

// GetOHLC method in API returns OHLC data and error(if error occured)
func (a *API) GetOHLC(ProductCode string, params string) (OHLC, error) {
	crypto := "https://api.cryptowat.ch/"
	path := ""
	var o OHLC

	//TODO: エイリアスじゃなくプロダクトコードに対応させる．MAT1WEEK, MAT2WEEKのプロダクトコードは自動生成させる.
	switch ProductCode {
	case "FX_BTC_JPY":
		path = "markets/bitflyer/btcfxjpy/ohlc"
	case "BTC_JPY":
		path = "markets/bitflyer/btcjpy/ohlc"
	case "ETH_BTC":
		path = "markets/bitflyer/ethbtc/ohlc"
	case "BCH_BTC":
		path = "markets/bitflyer/bchbtc/ohlc"
	case "BTCJPY_MAT1WK":
		path = "markets/bitflyer/btcjpy/ohlc"
	case "BTCJPY_MAT2WK":
		path = "markets/bitflyer/btcjpy/ohlc"
	default:
		return o, fmt.Errorf("not found product_code: %s", ProductCode)
	}

	byteArray := a.Request(crypto, path, "GET", params)
	err := json.Unmarshal(byteArray, &o)

	if err != nil {
		log.Fatal(err)
	}

	return o, nil
}

// Market struct represents markets from bitflyer
type Market struct {
	ProductCode string `json:"product_code"`
	Alias       string `json:"alias"`
}

// GetMarkets method in API
func (a *API) GetMarkets() ([]Market, error) {
	bitflyer := "https://api.bitflyer.jp/v1/"
	path := "markets"

	var m []Market
	byteArray := a.Request(bitflyer, path, "GET", "")
	err := json.Unmarshal(byteArray, &m)

	return m, err
}

type Boards struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

type Order struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

func (a *API) GetBoard(ProductCode string) (Boards, error) {
	bitflyer := "https://api.bitflyer.jp/v1/"
	path := "board"
	params := "?product_code=" + ProductCode

	var b Boards
	byteArray := a.Request(bitflyer, path, "GET", params)
	err := json.Unmarshal(byteArray, &b)

	return b, err
}
