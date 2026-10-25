package main

import "testing"

func TestTransfer(t *testing.T) {
	a := Account{Balance: 1000}
	b := Account{Balance: 300}

	err := a.Transfer(&b, 400)

	if err != nil {
		t.Fatalf("не ожидали ошибку, получили: %v", err)
	}

	if a.Balance != 600 {
		t.Errorf("баланс отправителя: ожидали 600, получили %d", a.Balance)
	}

	if b.Balance != 700 {
		t.Errorf("баланс получателя: ожидали 700, получили %d", b.Balance)
	}
}

func TestTransferNegativeSum(t *testing.T) {
	a := Account{Balance: 1000}
	b := Account{Balance: 300}

	err := a.Transfer(&b, -50)

	if err == nil {
		t.Fatalf("ожидали ошибку, но ее не было")
	}

	if a.Balance != 1000 {
		t.Errorf("баланс отправителя: ожидали 1000, получили %d", a.Balance)

	}

	if b.Balance != 300 {
		t.Errorf("баланс получателся: ожидали 300, получили %d", b.Balance)
	}

}

func TestTransferZeroSum(t *testing.T) {
	a := Account{Balance: 1000}
	b := Account{Balance: 300}

	err := a.Transfer(&b, 0)

	if err == nil {
		t.Fatalf("ожидали ошибку, но ее не было")
	}

	if a.Balance != 1000 {
		t.Errorf("баланс отправителя: ожидали 1000, получили %d", a.Balance)

	}

	if b.Balance != 300 {
		t.Errorf("баланс получателся: ожидали 300, получили %d", b.Balance)
	}
}

func TestTransferInsufficientFunds(t *testing.T) {
	a := Account{Balance: 1000}
	b := Account{Balance: 300}

	err := a.Transfer(&b, 1500)

	if err == nil {
		t.Fatalf("ожидали ошибку, но ее не было")

	}

	if a.Balance != 1000 {
		t.Errorf("баланс отправителя: ожидали 1000, получили %d", a.Balance)

	}

	if b.Balance != 300 {
		t.Errorf("баланс получателся: ожидали 300, получили %d", b.Balance)
	}
}

func TestTransferZeroBalance(t *testing.T) {
	a := Account{Balance: 0}
	b := Account{Balance: 300}

	err := a.Transfer(&b, 100)

	if err == nil {
		t.Fatalf("ожидали ошибку, но ее не было")

	}

	if a.Balance != 0 {
		t.Errorf("баланс отправителя: ожидали 0, получили %d", a.Balance)

	}

	if b.Balance != 300 {
		t.Errorf("баланс получателся: ожидали 300, получили %d", b.Balance)
	}
}
