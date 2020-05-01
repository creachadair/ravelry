// Package ravelry is a client for the Ravelry API.
//
// Ravelry: https://www.ravelry.com
//
// API documentation: https://www.ravelry.com/api
package ravelry

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPClient is the interface required to call the Ravelry API service.
// It is satisfied by the net/http Client type.
// Requests are created using http.NewRequestWithContext.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// API is the base URL of the production API endpoint.
const API = "https://api.ravelry.com"

type Client struct {
	// The base URL of the API endpoint. If empty, use the production API.
	API string

	// The HTTP client used to call the API. If nil, use http.DefaultClient.
	HTTPClient HTTPClient
}

type authTokenKey struct{}

type authToken struct {
	Type  string // e.g., Basic, Bearer
	Token string // fully encoded for transport
}

// WithBasicAuth adds basic authentication credentials to the provided context.
func WithBasicAuth(ctx context.Context, user, token string) context.Context {
	auth := user + ":" + token
	return context.WithValue(ctx, authTokenKey{}, authToken{
		Type:  "Basic",
		Token: base64.StdEncoding.EncodeToString([]byte(auth)),
	})
}

// TimeFormat is the timestamp format used by the Ravelry API.
const TimeFormat = "2006/01/02 15:04:05 -0700"

// A Timestamp represents a timestamp as encoded by the Ravelry API.
type Timestamp time.Time

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	var stamp string
	if err := json.Unmarshal(data, &stamp); err != nil {
		return fmt.Errorf("incorrect timestamp format: %v", err)
	}
	v, err := time.Parse(TimeFormat, stamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %v", err)
	}
	*ts = Timestamp(v)
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(ts).Format(TimeFormat) + `"`), nil
}

// PageInfo records the request  parameters controlling pagination.
type PageInfo struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// A Paginator is sent in replies large enough to require pagination.
type Paginator struct {
	PageCount int `json:"page_count"`
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Results   int `json:"results"`
	LastPage  int `json:"last_page"`
}
