package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.coingecko.com/api/v3/"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetPrice(coinID, currency string) (float64, error) {
	url := fmt.Sprint("%s/simple/price?ids=%s&vs_currencies=%s", baseURL, coinID, currency)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API error: %s", resp.Status)
	}
	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("Failed to decode response: %w", err)
	}

	if price, ok := result[coinID][currency]; ok {
		return price, nil
	}

	return 0, fmt.Errorf("price not found for %s/%s", coinID, currency)

}
