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

type ohlcResult struct {
	CloseTime  *string
	OpenPrice  *float64
	HighPrice  *float64
	LowPrice   *float64
	ClosePrice *float64
	Volume     *float64
}
type ohlcAllowance struct {
	Cost      int64 `json:"cost"`
	Remaining int64 `json:"remaining"`
}

// OHLC struct from cryptwatch
type OHLC struct {
	Results   []ohlcResult
	Allowance ohlcAllowance
}

// Market struct represents markets from bitflyer
type Market struct {
	ProductCode string `json:"product_code"`
	Alias       string `json:"alias"`
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
func (a *API) Request(url string, endPoint string, method string, params string) string {
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

	return string(byteArray)
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

	stringArray := a.Request(crypto, path, "GET", params)
	err := json.Unmarshal([]byte(stringArray), &o)

	if err != nil {
		log.Fatal(err)
	}

	return o, nil
}

// GetMarkets method in API
func (a *API) GetMarkets() []Market {
	bitflyer := "https://api.bitflyer.jp/v1/"
	path := "markets"

	var m []Market
	stringArray := a.Request(bitflyer, path, "GET", "")
	err := json.Unmarshal([]byte(stringArray), &m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}
