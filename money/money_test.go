package money

import (
	"testing"
)

func assertEquals(t *testing.T, target interface{}, values ...interface{}) {
	for _, val := range values {
		if target != val {
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
		assertEquals(t, *NewDollar(10), *five.Times(2))
		assertEquals(t, *NewDollar(15), *five.Times(3))
	})

	t.Run("equality", func(t *testing.T) {
		assertTrue(t, Equals(NewDollar(5), NewDollar(5)))
		assertFalse(t, Equals(NewDollar(5), NewDollar(6)))
		assertTrue(t, Equals(NewFranc(5), NewFranc(5)))
		assertFalse(t, Equals(NewFranc(5), NewFranc(6)))
		assertFalse(t, Equals(NewFranc(5), NewDollar(5)))
	})

	t.Run("Franc multiplication", func(t *testing.T) {
		five := NewFranc(5)
		assertEquals(t, *NewFranc(10), *five.Times(2))
		assertEquals(t, *NewFranc(15), *five.Times(3))
	})
}
