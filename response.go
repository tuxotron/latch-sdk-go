package latch

import "encoding/json"

type LatchResponse struct {
	Data  json.RawMessage
	Error latchErrorResponse
}

type latchErrorResponse struct {
	Code int
	Message string
}
