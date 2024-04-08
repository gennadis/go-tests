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

}
