package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(bufio.NewReader(os.Stdin), os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

func run(in io.Reader, out io.Writer) error {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return fmt.Errorf("введите количество чисел: %w", err)
	}
	if n <= 0 {
		return fmt.Errorf("количество чисел должно быть положительным")
	}
	var minimum, maximum float64
	for i := 0; i < n; i++ {
		var value float64
		if _, err := fmt.Fscan(in, &value); err != nil {
			return fmt.Errorf("число %d: %w", i+1, err)
		}
		if i == 0 || value < minimum {
			minimum = value
		}
		if i == 0 || value > maximum {
			maximum = value
		}
	}
	fmt.Fprintf(out, "Максимум: %g\nМинимум: %g\n", maximum, minimum)
	return nil
}
