package pointerr

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Value: 10}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{Value: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsuffientFunds)
		assertBalance(t, wallet, startingBalance)

	})
}
