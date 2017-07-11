package latch

import "encoding/json"

type LatchResponse struct {
	Data json.RawMessage
	Error json.RawMessage
}
