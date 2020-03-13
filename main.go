package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	//BASEURL := "https://api.binance.com"
	//ping(BASEURL)
	fmt.Println(TestValid())

}
