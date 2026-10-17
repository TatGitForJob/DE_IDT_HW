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
	if n < 0 {
		return fmt.Errorf("количество чисел не может быть отрицательным")
	}
	positive, negative, zeros := 0, 0, 0
	for i := 0; i < n; i++ {
		var value float64
		if _, err := fmt.Fscan(in, &value); err != nil {
			return fmt.Errorf("число %d: %w", i+1, err)
		}
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zeros++
		}
	}
	fmt.Fprintf(out, "Положительных: %d\nОтрицательных: %d\nНулей: %d\n", positive, negative, zeros)
	return nil
}
