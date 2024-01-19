package tuple

import (
	"testing"

	testutilities "github.com/jebschiefer/ray-tracer-challenge/src/test-utilities"
)

func TestIsPoint(t *testing.T) {
	tuple := Tuple{
		x: 4.3,
		y: -4.2,
		z: 3.1,
		w: 1.0,
	}

	testutilities.CompareFloats(t, tuple.x, 4.3)
	testutilities.CompareFloats(t, tuple.y, -4.2)
	testutilities.CompareFloats(t, tuple.z, 3.1)
	testutilities.CompareFloats(t, tuple.w, 1.0)

	testutilities.CompareBools(t, tuple.isPoint(), true)
	testutilities.CompareBools(t, tuple.isVector(), false)
}

func TestIsVector(t *testing.T) {
	tuple := Tuple{
		x: 4.3,
		y: -4.2,
		z: 3.1,
		w: 0.0,
	}

	testutilities.CompareFloats(t, tuple.x, 4.3)
	testutilities.CompareFloats(t, tuple.y, -4.2)
	testutilities.CompareFloats(t, tuple.z, 3.1)
	testutilities.CompareFloats(t, tuple.w, 0.0)

	testutilities.CompareBools(t, tuple.isPoint(), false)
	testutilities.CompareBools(t, tuple.isVector(), true)
}

func TestCreatePoint(t *testing.T) {
	tuple := Point(4.0, -4.0, 3.0)
	testutilities.CompareBools(t, tuple.isPoint(), true)
	testutilities.CompareBools(t, tuple.isVector(), false)
}

func TestCreateVector(t *testing.T) {
	tuple := Vector(4.0, -4.0, 3.0)
	testutilities.CompareBools(t, tuple.isPoint(), false)
	testutilities.CompareBools(t, tuple.isVector(), true)
}
