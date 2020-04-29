// Package ravelry is a client for the Ravelry API.
//
// Ravelry: https://www.ravelry.com
//
// API documentation: https://www.ravelry.com/api
package ravelry

import (
	"context"
	"encoding/base64"
	"net/http"
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
