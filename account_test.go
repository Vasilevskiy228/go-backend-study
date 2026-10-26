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

type testCase struct {
	name        string
	fromBalance int
	toBalance   int
	sum         int
	wantFrom    int
	wantTo      int
	wantErr     bool
}

func TestTransferTable(t *testing.T) {
	tests := []testCase{
		{
			name:        "тест перевода",
			fromBalance: 1000,
			toBalance:   300,
			sum:         400,
			wantFrom:    600,
			wantTo:      700,
			wantErr:     false,
		},
		{
			name:        "тест с отрицательной суммой",
			fromBalance: 1000,
			toBalance:   300,
			sum:         -50,
			wantFrom:    1000,
			wantTo:      300,
			wantErr:     true,
		},
		{
			name:        "тест нулевой суммы",
			fromBalance: 1000,
			toBalance:   300,
			sum:         0,
			wantFrom:    1000,
			wantTo:      300,
			wantErr:     true,
		},
		{
			name:        "тест перевод больше баланса",
			fromBalance: 1000,
			toBalance:   300,
			sum:         1500,
			wantFrom:    1000,
			wantTo:      300,
			wantErr:     true,
		},
		{
			name:        "тест нулевой баланс",
			fromBalance: 0,
			toBalance:   300,
			sum:         100,
			wantFrom:    0,
			wantTo:      300,
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := Account{Balance: tc.fromBalance}
			b := Account{Balance: tc.toBalance}

			err := a.Transfer(&b, tc.sum)

			gotErr := err != nil

			if gotErr != tc.wantErr {
				t.Fatalf("ожидали оишбку: %v, получили: %v", tc.wantErr, gotErr)
			}

			if a.Balance != tc.wantFrom {
				t.Errorf("баланс отправителя: ожидали %d, получили: %d", tc.fromBalance, a.Balance)

			}

			if b.Balance != tc.wantTo {
				t.Errorf("баланс получателя: ожидали %d,получили: %d", tc.toBalance, b.Balance)
			}
		})
	}
}
