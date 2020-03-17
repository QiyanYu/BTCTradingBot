package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"./signature"
)

func ping(BASEURL string) {
	pingURL := "/api/v3/ping"
	resp, err := http.Get(BASEURL + pingURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	context := "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"
	signature.GetSignature(context)

}
