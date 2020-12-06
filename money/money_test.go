package money

import (
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, target interface{}, values ...interface{}) {
	for _, val := range values {
		if !reflect.DeepEqual(target, val) {
			t.Fatalf("%v != %v", target, val)
		}
	}
}

func assertTrue(t *testing.T, eval bool) {
	if !eval {
		t.Fatalf("not true")
	}
}

func assertFalse(t *testing.T, eval bool) {
	if eval {
		t.Fatalf("not false")
	}
}

func TestMoney(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		five := NewDollar(5)
		assertEquals(t, NewDollar(10), five.Times(2))
		assertEquals(t, NewDollar(15), five.Times(3))
	})

	t.Run("equality", func(t *testing.T) {
		assertTrue(t, NewDollar(5).Equals(NewDollar(5)))
		assertFalse(t, NewDollar(5).Equals(NewDollar(6)))
		assertFalse(t, NewFranc(5).Equals(NewDollar(5)))
	})

	t.Run("currency", func(t *testing.T) {
		assertEquals(t, "USD", NewDollar(1).Currency())
		assertEquals(t, "CHF", NewFranc(1).Currency())
	})

	t.Run("simple addition", func(t *testing.T) {
		five := NewDollar(5)
		sum := five.Plus(five)
		bank := NewBank()
		reduced := bank.Reduce(sum, "USD")
		assertEquals(t, NewDollar(10), reduced)
	})

	t.Run("Plus returns Sum", func(t *testing.T) {
		five := NewDollar(5)
		result := five.Plus(five)
		sum := result.(*Sum)
		assertEquals(t, five, sum.Augend)
		assertEquals(t, five, sum.Addend)
	})

	t.Run("reduce Sum", func(t *testing.T) {
		sum := NewSum(NewDollar(3), NewDollar(4))
		bank := NewBank()
		result := bank.Reduce(sum, "USD")
		assertEquals(t, NewDollar(7), result)
	})

	t.Run("reduce Money", func(t *testing.T) {
		bank := NewBank()
		result := bank.Reduce(NewDollar(1), "USD")
		assertEquals(t, NewDollar(1), result)
	})

	t.Run("reduce Money different currency", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(NewFranc(2), "USD")
		assertEquals(t, NewDollar(1), result)
	})

	t.Run("identity rate", func(t *testing.T) {
		assertEquals(t, 1, NewBank().Rate("USD", "USD"))
	})

	t.Run("mixed addition", func(t *testing.T) {
		fiveBucks := (Expression)(NewDollar(5))
		tenFrancs := (Expression)(NewFranc(10))
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
		assertEquals(t, NewDollar(10), result)
	})
}
