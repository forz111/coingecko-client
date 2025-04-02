package coingecko

type Coin struct {
	ID     string
	Name   string
	Symbol string
}

type MarketData struct {
	CurrentPrice map[string]float64 `json:"current_price"`
	MarketCap    map[string]float64 `json:"market_cap"`
}
