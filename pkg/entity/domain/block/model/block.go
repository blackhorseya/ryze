package model

import (
	"encoding/json"
	"strconv"
	"time"
)

func (x *Block) MarshalJSON() ([]byte, error) {
	type Alias Block

	return json.Marshal(&struct {
		*Alias
		Number    string `json:"number,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	}{
		Alias:     (*Alias)(x),
		Number:    strconv.Itoa(int(x.Number)),
		Timestamp: x.Timestamp.AsTime().UTC().Format(time.RFC3339),
	})
}
