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
		{"mixed", "5 -1 0 2 -3 4", "Положительных: 2\nОтрицательных: 2\nНулей: 1\n", false},
		{"empty", "0", "Положительных: 0\nОтрицательных: 0\nНулей: 0\n", false},
		{"negative_count", "-1", "", true},
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
