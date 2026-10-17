package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name, input, want string
		wantErr           bool
	}{
		{"view", "100 1", "Баланс: 100.00\n", false},
		{"deposit", "100 2 25.5", "Баланс: 125.50\n", false},
		{"withdraw", "100 3 40", "Баланс: 60.00\n", false},
		{"all", "100 3 100", "Баланс: 0.00\n", false},
		{"insufficient", "100 3 101", "", true},
		{"negative_balance", "-1 1", "", true},
		{"negative_amount", "100 2 -10", "", true},
		{"zero_amount", "100 3 0", "", true},
		{"bad_action", "100 4", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			err := run(strings.NewReader(tt.input), &out)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr = %v", err, tt.wantErr)
			}
			if out.String() != tt.want {
				t.Fatalf("output = %q, want %q", out.String(), tt.want)
			}
		})
	}
}
