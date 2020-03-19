package request

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.binance.com"

//Ping for ping the API service
func Ping() string {
	url := baseURL + "/api/v3/time"
	return sendNoneRequest(url)
}

//GetCurrentAvgPrice for get current average price
func GetCurrentAvgPrice(symbol string) string {
	url := baseURL + "/api/v3/avgPrice?symbol=" + symbol
	return sendNoneRequest(url)
}

//GetSymbolPriceTicker for get the latest price of a cryptocurrency
func GetSymbolPriceTicker(symbol string) string {
	url := baseURL + "/api/v3/ticker/price?symbol=" + symbol
	return sendNoneRequest(url)
}

//These are five type of security requests: 1)NONE 2)TRADE 3)USER_DATA 4)USER_STREAM 5)MARKET_DATA
//sendNoneRequest for sending None type of request
func sendNoneRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	return string(body)
}
