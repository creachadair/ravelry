// Package types defines the request and response types in the Ravelry API.
package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimeFormat is the timestamp format used by the Ravelry API.
const TimeFormat = "2006/01/02 15:04:05 -0700"

// A Date represents a timestamp as encoded by the Ravelry API.
type Date time.Time

// UnmarshalJSON implements json.Unmarshaler to decode a date from a string.
func (d *Date) UnmarshalJSON(data []byte) error {
	var stamp string
	if err := json.Unmarshal(data, &stamp); err != nil {
		return fmt.Errorf("incorrect timestamp format: %v", err)
	}
	v, err := time.Parse(TimeFormat, stamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %v", err)
	}
	*d = Date(v)
	return nil
}

// MarshalJSON implements json.Marshaler to encode a date as a string.
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format(TimeFormat) + `"`), nil
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
