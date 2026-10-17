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
		{"negative", "3 -9 -2 -7", "Максимум: -2\nМинимум: -9\n", false},
		{"one", "1 2.5", "Максимум: 2.5\nМинимум: 2.5\n", false},
		{"empty", "0", "", true},
		{"missing", "2 1", "", true},
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
