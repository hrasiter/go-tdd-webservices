package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Wallet Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, Bitcoin(10), wallet)

	})

	t.Run("Wallet Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, Bitcoin(10), wallet)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, startingBalance, wallet)
		assertError(t, err, ErrInsufficientFund)
	})

}

func assertBalance(t testing.TB, want Bitcoin, w Wallet) {
	t.Helper()
	got := w.Balance()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("want Error, but got nil!")
	}

	if err != want {
		t.Errorf("want %q, got: %q", want, err.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("want nil, but got error %q", err)
	}
}
