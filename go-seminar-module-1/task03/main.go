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
	if n <= 0 {
		return fmt.Errorf("N должно быть положительным")
	}
	var sum float64
	for i := 0; i < n; i++ {
		sum += float64(i + 1)
	}
	fmt.Fprintf(out, "Сумма: %g\nСреднее: %g\n", sum, sum/float64(n))
	return nil
}
