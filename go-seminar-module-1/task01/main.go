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
	var a, b float64
	var operation string
	if _, err := fmt.Fscan(in, &a, &b, &operation); err != nil {
		return fmt.Errorf("введите два числа и операцию: %w", err)
	}
	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return fmt.Errorf("неизвестная операция: %s", operation)
	}
	fmt.Fprintf(out, "Результат: %g\n", result)
	return nil
}
