package test

import "testing"

func TestCompare(t *testing.T) {
	compared := Compare("a", "b")
	expected := -1

	if compared != expected {
		t.Errorf("expected %q but got %q", expected, compared)
	}

	compared = Compare("a", "a")
	expected = 0

	if compared != expected {
		t.Errorf("expected %q but got %q", expected, compared)
	}

	compared = Compare("b", "a")
	expected = 1

	if compared != expected {
		t.Errorf("expected %q but got %q", expected, compared)
	}
}
