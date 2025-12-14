package request

import (
	"errors"
	"strings"
	"testing"
)

func floatPtr(val float64) *float64 {
	return &val
}

func TestCreateReqTableDriven(t *testing.T) {
	var tests = []struct {
		name    string
		input   []byte
		wantReq *Request
		wantErr string
	}{
		// --- Success  ---
		{
			name:    "Success Case",
			input:   []byte(`{"method":"isPrime","number":123}`),
			wantReq: &Request{Method: "isPrime", Number: floatPtr(123)},
			wantErr: "",
		},
		{
			name:    "Success Case with 0",
			input:   []byte(`{"method":"isPrime","number":0}`),
			wantReq: &Request{Method: "isPrime", Number: floatPtr(0)},
			wantErr: "",
		},
		{
			name:    "Success Case with extra field",
			input:   []byte(`{"method":"isPrime","number":123, "ignored":2}`),
			wantReq: &Request{Method: "isPrime", Number: floatPtr(123)},
			wantErr: "",
		},
		{
			name:    "Success Case with Float",
			input:   []byte(`{"method":"isPrime","number":123.456}`),
			wantReq: &Request{Method: "isPrime", Number: floatPtr(123.456)},
			wantErr: "",
		},
		// --- Failed ---
		{
			name:    "Missing Method",
			input:   []byte(`{"number":123}`),
			wantReq: nil,
			wantErr: "invalid request",
		},
		{
			name:    "Incorrect Method",
			input:   []byte(`{"method":"isNotPrime","number":123}`),
			wantReq: nil,
			wantErr: "invalid request",
		},
		{
			name:    "Malformed JSON - Type Mismatch",
			input:   []byte(`{"method":"isPrime","number":"123"}`),
			wantReq: nil,
			wantErr: "invalid request",
		},
		{
			name:    "Invalid JSON",
			input:   []byte(`not valid json`),
			wantReq: nil,
			wantErr: "invalid request",
		},
		{
			name:    "Missing Number field (nil pointer check)",
			input:   []byte(`{"method":"isPrime"}`),
			wantReq: nil,
			wantErr: "invalid request",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotReq, err := CreateReq(tt.input)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("Expected error containing '%s', but got nil.", tt.wantErr)
				}
				if !errors.Is(err, errors.New(tt.wantErr)) {
					if !strings.Contains(err.Error(), tt.wantErr) {
						t.Errorf("Error mismatch. Got error: '%v', Expected error to contain: '%s'", err, tt.wantErr)
					}
				}

				if gotReq != nil {
					t.Errorf("Expected nil request on error, but got: %+v", gotReq)
				}
				return
			}

			if gotReq.Method != tt.wantReq.Method {
				t.Errorf("Method mismatch. Got '%s', want '%s'", gotReq.Method, tt.wantReq.Method)
			}

			if *gotReq.Number != *tt.wantReq.Number {
				t.Errorf("Number mismatch. Got '%f', want '%f'", *gotReq.Number, *tt.wantReq.Number)
			}
		})
	}
}
