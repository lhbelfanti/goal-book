// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 262.

// Package bank provides a concurrency-safe bank with one account.
package bank_test

import (
	"sync"
	"testing"

	bank "gopl.io/ch9/bank1"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
