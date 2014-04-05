package main

import "testing"

func TestCalcAbsDiff(t *testing.T) {
	const x, y, expected uint = 3, 20, 17
	actual := CalcAbsDiff(x, y)

	if actual != expected {
		t.Errorf("CalcAbsDiff(%v, %v) = %v, want %v", x, y, actual, expected)
	}
}
