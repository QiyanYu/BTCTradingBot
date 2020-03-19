package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"./request"
	"./signature"
)

var baseURL = "https://api.binance.com"

func getTime() string {
	// timenow := time.Now().Unix()
	time := time.Now().UnixNano() / 1e6
	fmt.Println(time)
	fmt.Println(strconv.FormatInt(time, 10))
	return strconv.FormatInt(time, 10)
}

func main() {

	c := signature.Conf{}
	c.GetConf()

	url := baseURL + "/sapi/v1/accountSnapshot?"
	time := getTime()
	fmt.Println(time)
	//since the server unix time is faster than mine, increased the recvWindow
	allCoinContext := "type=SPOT&recvWindow=50000&limit=5&timestamp=" + time
	urlAllCoin := url + allCoinContext + "&signature=" + c.GetSignature(allCoinContext)
	fmt.Println(urlAllCoin)
	req, err := http.NewRequest("GET", urlAllCoin, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-MBX-APIKEY", c.GetAPIKey())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		log.Fatal("Error reading body, ", err)
	}

	fmt.Println(string(body))

	request.Ping()
	fmt.Println(request.GetSymbolPriceTicker("BTCUSDT"))
	request.WSLiveSubscribe()
}
