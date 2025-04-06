//go:build ignore

// Пакетный запрос цен

package main

import (
	"fmt"
	"log"

	"github.com/forz111/coingecko-client"
)

func main() {
	client := coingecko.New()
	coins := []string{"bitcoin", "ethereum", "solana"}

	for _, coin := range coins {
		price, err := client.GetPriceWithRetry(coin, "usd", 3)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s: $%.2f\n", coin, price)
	}
}
