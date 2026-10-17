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
	var sum float64
	var minimum, maximum int64
	even, odd := 0, 0
	for i := 0; i < n; i++ {
		var value int64
		if _, err := fmt.Fscan(in, &value); err != nil {
			return fmt.Errorf("число %d: %w", i+1, err)
		}
		sum += float64(value)
		if i == 0 || value < minimum {
			minimum = value
		}
		if i == 0 || value > maximum {
			maximum = value
		}
		if value%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Fprintf(out, "Сумма: %g\nСреднее: %g\nМаксимум: %d\nМинимум: %d\nЧётных: %d\nНечётных: %d\n", sum, sum/float64(n), maximum, minimum, even, odd)
	return nil
}
