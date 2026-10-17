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
		{"mixed", "5 -3 0 2 5 -4", "Сумма: 0\nСреднее: 0\nМаксимум: 5\nМинимум: -4\nЧётных: 3\nНечётных: 2\n", false},
		{"negative", "3 -7 -2 -3", "Сумма: -12\nСреднее: -4\nМаксимум: -2\nМинимум: -7\nЧётных: 1\nНечётных: 2\n", false},
		{"one", "1 5", "Сумма: 5\nСреднее: 5\nМаксимум: 5\nМинимум: 5\nЧётных: 0\nНечётных: 1\n", false},
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
