package testutilities

import "testing"

func CompareBools(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func CompareFloats(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
