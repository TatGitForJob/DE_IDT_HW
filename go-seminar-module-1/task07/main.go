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
		return fmt.Errorf("введите N: %w", err)
	}
	if n < 0 {
		return fmt.Errorf("N не может быть отрицательным")
	}
	var sum int64
	for i := 0; i < n; i++ {
		value := i + 1
		if value%2 == 0 && value%3 == 0 {
			sum += int64(value)
		}
	}
	fmt.Fprintf(out, "Сумма: %d\n", sum)
	return nil
}
