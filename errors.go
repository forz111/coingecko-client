package coingecko

import (
	"errors"
	"fmt"
)

var (
	ErrAPIRequestFailed = errors.New("API request failed")
	ErrDecodingFailed = errors.New("decoding failed")
	ErrRateLimit = errors.New("rate limit exceeded")

	errCoinNotFound = errors.New("coin not found")
	errInvalidCurrency = errors.New("invalid currency")

)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	RequestURL string `json:"request_url,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

func IsCoinNotFound(err error) bool {
	return errors.Is(err, errCoinNotFound)
}

func IsInvalidCurrency(err error) bool {
	return errors.Is(err, errInvalidCurrency)
}
