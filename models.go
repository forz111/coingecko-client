package coingecko

import "time"

type Coin struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Symbol    string         `json:"symbol"`
	Platforms []CoinPlatform `json:"platforms,omitempty"`
}

type CoinPlatform struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type MarketData struct {
	CurrentPrice   map[string]float64 `json:"current_price"`
	MarketCap      map[string]float64 `json:"market_cap"`
	PriceChange24h float64            `json:price_change_24h`
	LastUpdated     time.Time `json:"last_updated"`
}

type PriceResponse struct {
	CoinID  string `json:"coin_id"`
	Prices map[string]float64 `json:"prices"`
	Timestamp time.Time `json:"timestamp"`
}
