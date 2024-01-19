package tuple

import "github.com/jebschiefer/ray-tracer-challenge/src/utilities"

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (a Tuple) equals(b Tuple) bool {
	return utilities.FloatsEqual(a.x, b.x) &&
		utilities.FloatsEqual(a.y, b.y) &&
		utilities.FloatsEqual(a.z, b.z) &&
		utilities.FloatsEqual(a.w, b.w)
}

func (tuple Tuple) isPoint() bool {
	return tuple.w == 1
}

func (tuple Tuple) isVector() bool {
	return tuple.w == 0
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
