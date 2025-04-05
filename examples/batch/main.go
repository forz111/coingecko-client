//go:build ignore

// Пакетный запрос цен

package main

import (
	"fmt"

	"github.com/forz111/coingecko-client"
)

func main() {
	client := coingecko.New()
	coins := []string{"bitcoin", "ethereum", "solana"}

	for _, coin := range coins {
		price, _ := client.GetPrice(coin, "usd")
		fmt.Printf("%s: $%.2f\n", coin, price)
	}
}
