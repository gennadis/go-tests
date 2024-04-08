package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Tugbacoin(10))
	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(5)
		assertNoError(t, err)
		assertBalance(t, wallet, Tugbacoin(5))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Tugbacoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(50)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Tugbacoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
