package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	}

	assertError := func(t testing.TB, err error, expected string) {
		t.Helper()

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err.Error() != expected {
			t.Errorf("expected %q error message, got %q", expected, err)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
