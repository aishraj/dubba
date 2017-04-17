package dubba

import "testing"

func TestGcounterValue(t *testing.T) {
	gCounter := NewGCounter()
	gCounter.Increment()
	currentValue := gCounter.Value()
	if currentValue != 1 {
		t.Errorf("Invalid value for g counter")
	}
}

func TestGCounterIncrement(t *testing.T) {
	gCounter := NewGCounter()
	gCounter.Increment()
	gCounter.IncrementNode("x", 4)
	gCounter.IncrementNode("a", 3)
	gCounter.IncrementNode("4", 1)
	if gCounter.Value() != 9 {
		t.Errorf("Invalid valid for G-Counter")
	}
}

func TestGCounterMerge(t *testing.T) {
	a, b := NewGCounter(), NewGCounter()
	a.Increment()
	b.IncrementNode("1", 9)
	a.IncrementNode("2", 91)
	ret := a.Merge(*b)
	if ret.Value() != 100 {
		t.Errorf("Invalid value for G-Counter. Merge failed")
	}
}

func TestGCounterEqual(t *testing.T) {
	a, b := NewGCounter(), NewGCounter()
	a.Increment()
	b.IncrementNode("1", 2)
	a.IncrementNode("1", 1)
	ret := a.IsEqualTo(*b)
	if ret != true {
		t.Errorf("Invalid value for G-Counter. Merge failed")
	}
}
