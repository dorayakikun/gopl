package bank_test

import (
	"fmt"
	"testing"

	"z"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	var ok bool
	go func() {
		ok = bank.Withdraw(150)
		done <- struct{}{}
	}()

	<-done

	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if !ok {
		t.Errorf("ok = %t, want %t", ok, true)
	}

	go func() {
		ok = bank.Withdraw(300)
		done <- struct{}{}
	}()

	<-done

	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	if ok {
		t.Errorf("ok = %t, want %t", ok, false)
	}
}
