package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Ticker contains ticket attributes
type Ticker struct {
	Timestamp int64   `json:"timestamp,string"`
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Bid       float64 `json:"bid,string"`
	Vwap      float64 `json:"vwap,string"`
	Volume    float64 `json:"volume,string"`
	Low       float64 `json:"low,string"`
	Ask       float64 `json:"ask,string"`
	Open      float64 `json:"open,string"`
}

// GetTicker retrieves the ticket data for the supplied pair
func GetTicker(pair string) (Ticker, error) {
	var ticker Ticker

	url := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%v", pair)
	resp, err := http.Get(url)

	if err != nil {
		log.Println("GET:", err)
		return ticker, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("READ:", err)
		return ticker, err
	}

	err = json.Unmarshal(body, &ticker)

	if err != nil {
		log.Println("UNMARSHALL:", err, "(", string(body), ")")
		return ticker, err
	}

	return ticker, err
}
