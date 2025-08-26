package generics

import (
	"testing"

	"github.com/lin-br/go-lin/learn-go/19-generics/asserts"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		asserts.AssertEqual(t, 1, 1)
		asserts.AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		asserts.AssertEqual(t, "hello", "hello")
		asserts.AssertNotEqual(t, "hello", "Grace")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := NewStack[int]()

		asserts.AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		asserts.AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		asserts.AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		asserts.AssertEqual(t, value, 123)
		asserts.AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		asserts.AssertEqual(t, firstNum+secondNum, 3)
	})
}
