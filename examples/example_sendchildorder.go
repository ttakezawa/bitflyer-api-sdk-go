package main

import (
	"log"
	"os"

	"github.com/fgken/bitflyer-api-sdk-go/bitflyerclient"
)

func main() {
	apiKey := os.Getenv("BITFLYER_API_KEY")
	apiSecret := os.Getenv("BITFLYER_API_SECRET")

	bfclient, err := bitflyerclient.New(apiKey, apiSecret)
	if err != nil {
		log.Fatal("Falied to new bitflyerclient")
	}

	order := bitflyerclient.MARKET
	side := bitflyerclient.BUY
	resp, err := bfclient.SendChildOrder(order, side, 0.001, 0)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
