package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"./signature"
)

func ping() {
	BASEURL := "https://api.binance.com"
	// pingURL := "/api/v3/ping"
	// systemURL := "/wapi/v3/systemStatus.html"
	systemURL := "/api/v3/time"
	resp, err := http.Get(BASEURL + systemURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(time.Now().UnixNano() / 1e6)
}
func getTime() string {
	// timenow := time.Now().Unix()
	time := time.Now().UnixNano() / 1e6
	fmt.Println(time)
	fmt.Println(strconv.FormatInt(time, 10))
	return strconv.FormatInt(time, 10)
}

func main() {
	BASEURL := "https://api.binance.com"

	c := signature.Conf{}
	c.GetConf()

	url := BASEURL + "/sapi/v1/accountSnapshot?"
	time := getTime()
	fmt.Println(time)
	//since the server unix time is faster than mine, increased the recvWindow
	allCoinContext := "type=SPOT&recvWindow=25000&limit=5&timestamp=" + time
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

	ping()
}
