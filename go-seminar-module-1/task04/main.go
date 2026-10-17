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
		return fmt.Errorf("введите целое число: %w", err)
	}
	printTableTo(out, n)
	return nil
}

func printTable(n int) { printTableTo(os.Stdout, n) }
func printTableTo(out io.Writer, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(out, "%d * %d = %d\n", n, i, n*i)
	}
}
