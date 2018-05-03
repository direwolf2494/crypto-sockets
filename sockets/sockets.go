package sockets

import (
	"log"
	"time"

	"github.com/direwolf2494/crypto-sockets/crypto"
	"golang.org/x/net/websocket"
)

// RetrieveTickerInfo uses an endless loop to continually retrieve ticker prices
func RetrieveTickerInfo(pair string, tickers chan *crypto.Ticker) {
	t := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-t.C:
			log.Println("Tick")
			ticker, err := crypto.GetTicker(pair)

			if err != nil {
				log.Println("ERROR: ", err)
			}

			log.Println("Adding ticker to channel")
			tickers <- &ticker
			log.Println("Ticker Data added.")
		default:
			// nothing
		}
	}
}

// TickerSocket returns ticker data for bitcoin over a websocket
func TickerSocket(ws *websocket.Conn) {
	tickers := make(chan *crypto.Ticker)
	var ticker crypto.Ticker

	go RetrieveTickerInfo("btcusd", tickers)

	for {
		// allocate our container struct
		select {
		case t := <-tickers:
			ticker = crypto.Ticker(*t)
			// send the ticker data
			if err := websocket.JSON.Send(ws, ticker); err != nil {
				log.Println(err)
			}
		default:
			// ntn
		}
	}
}
