package tuple

import (
	"testing"

	testutilities "github.com/jebschiefer/ray-tracer-challenge/src/test-utilities"
)

func TestIsPoint(t *testing.T) {
	tuple := tuple{
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
	tuple := tuple{
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

func TestEquals(t *testing.T) {
	a := Point(4.0, -4.0, 3.0)
	b := Point(4.0, -4.0, 3.0)
	c := Point(4.0, -4.0, 5.0)

	testutilities.CompareBools(t, a.equals(b), true)
	testutilities.CompareBools(t, a.equals(c), false)
}

func TestAddTwoTuples(t *testing.T) {
	a := Point(3, -2, 5)
	b := Vector(-2, 3, 1)
	c := Point(1, 1, 6)

	sum := a.add(b)

	testutilities.CompareBools(t, sum.equals(c), true)
}

func TestSubtractTwoPoints(t *testing.T) {
	a := Point(3, 2, 1)
	b := Point(5, 6, 7)
	c := Vector(-2, -4, -6)

	difference := a.subtract(b)

	testutilities.CompareBools(t, difference.equals(c), true)
}

func TestSubtractVectorFromPoint(t *testing.T) {
	a := Point(3, 2, 1)
	b := Vector(5, 6, 7)
	c := Point(-2, -4, -6)

	difference := a.subtract(b)

	testutilities.CompareBools(t, difference.equals(c), true)
}

func TestSubtractTwoVectors(t *testing.T) {
	a := Vector(3, 2, 1)
	b := Vector(5, 6, 7)
	c := Vector(-2, -4, -6)

	difference := a.subtract(b)

	testutilities.CompareBools(t, difference.equals(c), true)
}
