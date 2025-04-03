package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.coingecko.com/api/v3/"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	httpClient *http.Client
	baseURL string
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: baseURL,
	}
}

func (c *Client) GetPrice(coinID, currency string) (float64, error) {
	if coinID == "" {
		return 0, fmt.Errorf("%w: coinID is empty", errCoinNotFound)
	}

	url := fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=%s", c.baseURL, coinID, currency)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrAPIRequestFailed, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		return 0, ErrRateLimit
	case http.StatusNotFound:
		return 0, fmt.Errorf("%w: %s", errCoinNotFound, coinID)
	case http.StatusOK:

	default:
		return 0, &APIError{
			StatusCode: resp.StatusCode,
			Message: "unexpected status",
		}
	}

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("%w: %v", ErrDecodingFailed, err)
	}

	coinData, ok := result[coinID]
	if !ok {
		return 0, fmt.Errorf("%w: %s", errCoinNotFound, coinID)
	}

	price, ok := coinData[currency]
    if !ok {
        return 0, fmt.Errorf("%w: %s", errInvalidCurrency, currency)
    }

	return price, nil

}
