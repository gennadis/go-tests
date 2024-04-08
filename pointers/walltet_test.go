package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Tugbacoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Withdraw(5)

		got := wallet.Balance()
		want := Tugbacoin(5)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

}
