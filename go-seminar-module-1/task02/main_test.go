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
		{"negative", "-15", "Отрицательное\nНечётное\nДелится на 3: true\nДелится на 5: true\nДелится на 3 и 5: true\n", false},
		{"zero", "0", "Ноль\nЧётное\nДелится на 3: true\nДелится на 5: true\nДелится на 3 и 5: true\n", false},
		{"positive", "2", "Положительное\nЧётное\nДелится на 3: false\nДелится на 5: false\nДелится на 3 и 5: false\n", false},
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
