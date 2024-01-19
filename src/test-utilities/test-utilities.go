package testutilities

import (
	"math"
	"testing"
)

const EPSILON = 1e-4

func CompareBools(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func CompareFloats(t *testing.T, got float64, want float64) {
	if math.Abs(got-want) > EPSILON {
		t.Errorf("got %f, want %f", got, want)
	}
}
