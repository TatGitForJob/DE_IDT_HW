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
	var balance float64
	var action int
	if _, err := fmt.Fscan(in, &balance, &action); err != nil {
		return fmt.Errorf("введите баланс и номер действия: %w", err)
	}
	if balance < 0 {
		return fmt.Errorf("баланс не может быть отрицательным")
	}
	switch action {
	case 1:
	case 2, 3:
		var amount float64
		if _, err := fmt.Fscan(in, &amount); err != nil {
			return fmt.Errorf("введите сумму: %w", err)
		}
		if amount <= 0 {
			return fmt.Errorf("сумма должна быть положительной")
		}
		if action == 2 {
			balance += amount
		} else {
			if amount > balance {
				return fmt.Errorf("недостаточно средств")
			}
			balance -= amount
		}
	default:
		return fmt.Errorf("неизвестное действие: %d", action)
	}
	fmt.Fprintf(out, "Баланс: %.2f\n", balance)
	return nil
}
