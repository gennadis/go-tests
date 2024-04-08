package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Tugbacoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Tugbacoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Withdraw(5)
		assertBalance(t, wallet, Tugbacoin(5))
	})
	t.Run("insufficient funds", func(t *testing.T) {
		startingBalance := Tugbacoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(50)
		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})

}
