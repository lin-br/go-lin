package generics_test

import (
	"strings"
	"testing"

	"github.com/lin-br/go-lin/generics"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = generics.Account{Name: "Riya", Balance: 100}
		chris = generics.Account{Name: "Chris", Balance: 75}
		adil  = generics.Account{Name: "Adil", Balance: 200}

		transactions = []generics.Transaction{
			generics.NewTransaction(chris, riya, 100),
			generics.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account generics.Account) float64 {
		return generics.NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := generics.Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})

	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := generics.Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}
