package main

import "testing"

func TestWallet(t *testing.T) {
	
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalence(t, wallet, Bitcoin(10))
	})
	
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalence(t, wallet, Bitcoin(10))
	})
	
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startinBalance := Bitcoin(20)
		wallet := Wallet{Bitcoin(startinBalance)}
		err := wallet.Withdraw(Bitcoin(100))
		
		assertError(t, err, ErrInsufficientFunds)
		assertBalence(t, wallet, startinBalance)
	})
}

func assertBalence(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("want an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

