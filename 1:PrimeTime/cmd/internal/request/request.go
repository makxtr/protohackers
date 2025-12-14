package request

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string
	Number *float64
}

func CreateReq(c []byte) (*Request, error) {
	var r Request
	err := json.Unmarshal(c, &r)

	if err != nil || r.Method != "isPrime" || r.Number == nil {
		return nil, fmt.Errorf("invalid request: %v", string(c))
	}

	return &r, nil
}
