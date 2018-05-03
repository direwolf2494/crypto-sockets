package main

import (
	"log"
	"net/http"

	"github.com/direwolf2494/crypto-sockets/sockets"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ticker", websocket.Handler(sockets.TickerSocket))

	log.Println("Server starting on: 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
