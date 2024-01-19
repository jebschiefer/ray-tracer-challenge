package testutilities

import (
	"testing"

	"github.com/jebschiefer/ray-tracer-challenge/src/utilities"
)

func CompareBools(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func CompareFloats(t *testing.T, got float64, want float64) {
	if !utilities.FloatsEqual(got, want) {
		t.Errorf("got %f, want %f", got, want)
	}
}
