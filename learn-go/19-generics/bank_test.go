package generics_test

import (
	"testing"

	"github.com/lin-br/go-lin/generics"
)

func TestBadBank(t *testing.T) {
	transactions := []generics.Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, generics.BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, generics.BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, generics.BalanceFor(transactions, "Adil"), -25)
}
