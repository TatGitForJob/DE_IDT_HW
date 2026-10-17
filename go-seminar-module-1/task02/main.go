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
	var n int64
	if _, err := fmt.Fscan(in, &n); err != nil {
		return fmt.Errorf("введите целое число: %w", err)
	}
	if n > 0 {
		fmt.Fprintln(out, "Положительное")
	} else if n < 0 {
		fmt.Fprintln(out, "Отрицательное")
	} else {
		fmt.Fprintln(out, "Ноль")
	}
	if n%2 == 0 {
		fmt.Fprintln(out, "Чётное")
	} else {
		fmt.Fprintln(out, "Нечётное")
	}
	fmt.Fprintf(out, "Делится на 3: %t\nДелится на 5: %t\nДелится на 3 и 5: %t\n", n%3 == 0, n%5 == 0, n%3 == 0 && n%5 == 0)
	return nil
}
