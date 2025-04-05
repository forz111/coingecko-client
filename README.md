# CoinGecko API Client
Go client for CoinGecko cryptocurrency API

## 📦 Basic Usage

```go
package main

import (
	"fmt"
	"github.com/yourname/coingecko-client/coingecko"
)

func main() {
	client := coingecko.New()
	price, err := client.GetPrice("bitcoin", "usd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("BTC price: $%.2f", price)
}
```
