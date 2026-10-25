package main

import (
	"errors"
	"fmt"
)

func main() {
	a := Account{
		Balance: 0,
	}

	b := Account{
		Balance: 300,
	}

	err := a.Transfer(&b, 1500)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println(a.Balance, b.Balance)

}

type Account struct {
	Balance int
}

func (a *Account) Transfer(b *Account, sum int) error {

	if sum <= 0 {
		err := errors.New("Перевод 0 или отрицательного числа")
		return err
	}

	if a.Balance <= 0 {
		err := errors.New("недостаточно средств или минусовой баланс")
		return err
	}

	if sum > a.Balance {
		err := errors.New("недостаточно средств для перевода")
		return err
	}

	a.Balance -= sum
	b.Balance += sum
	return nil
}
