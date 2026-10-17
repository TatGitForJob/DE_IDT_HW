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
		{"addition", "2 3 +", "Результат: 5\n", false},
		{"subtraction", "2 5 -", "Результат: -3\n", false},
		{"multiplication", "-2 1.5 *", "Результат: -3\n", false},
		{"division", "7 2 /", "Результат: 3.5\n", false},
		{"zero", "1 0 /", "", true},
		{"operator", "1 2 %", "", true},
		{"bad_input", "abc", "", true},
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
