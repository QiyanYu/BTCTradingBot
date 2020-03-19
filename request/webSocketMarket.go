package request

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

const baseWsURL = "wss://stream.binance.com:9443"

//WSLiveSubscribe for getting Web Socket live subscribe
func WSLiveSubscribe() {
	url := baseWsURL + "/ws/btcusdt@aggTrade"
	sub := `{"method": "SUBSCRIBE","params":["btcusdt@aggTrade"],"id": 1}`
	ws, err := websocket.Dial(url, "", baseURL)
	if err != nil {
		log.Fatalln("1", err)
	}
	if _, err := ws.Write([]byte(sub)); err != nil {
		log.Fatalln("2", err)
	}
	var msg = make([]byte, 512)
	var n int
	num := 1
	for {

		if n, err = ws.Read(msg); err != nil {
			log.Fatalln("3", err)
		}
		fmt.Println("Received: ", num, " ", string(msg[:n]))
		num++
		time.Sleep(10 * time.Second)
	}
	// if n, err = ws.Read(msg); err != nil {
	// 	log.Fatalln("3", err)
	// }
	// fmt.Println("Received: ", string(msg[:n]))
}
