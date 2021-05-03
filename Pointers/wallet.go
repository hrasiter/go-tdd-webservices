package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFund = errors.New("insufficient fund")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if withdraw > w.balance {
		return ErrInsufficientFund
	}
	w.balance -= withdraw
	return nil
}
