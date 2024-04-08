package pointers

import (
	"errors"
	"fmt"
)

type Tugbacoin int

func (t Tugbacoin) String() string {
	return fmt.Sprintf("%d TGBC", t)
}

type Wallet struct {
	balance Tugbacoin
}

func (w *Wallet) Deposit(amount Tugbacoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Tugbacoin) error {
	if w.balance < amount {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Tugbacoin {
	return w.balance
}
