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
