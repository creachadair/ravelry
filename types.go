package ravelry

import (
	"encoding/json"
	"fmt"
	"time"
)

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
