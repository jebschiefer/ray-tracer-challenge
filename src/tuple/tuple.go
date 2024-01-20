package tuple

import "github.com/jebschiefer/ray-tracer-challenge/src/utilities"

type tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func Point(x, y, z float64) tuple {
	return tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) tuple {
	return tuple{x, y, z, 0.0}
}

func (a tuple) equals(b tuple) bool {
	return utilities.FloatsEqual(a.x, b.x) &&
		utilities.FloatsEqual(a.y, b.y) &&
		utilities.FloatsEqual(a.z, b.z) &&
		utilities.FloatsEqual(a.w, b.w)
}

func (tuple tuple) isPoint() bool {
	return tuple.w == 1
}

func (tuple tuple) isVector() bool {
	return tuple.w == 0
}

func (a tuple) add(b tuple) tuple {
	return tuple{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
		a.w + b.w,
	}
}

func (a tuple) subtract(b tuple) tuple {
	return tuple{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
		a.w - b.w,
	}
}

func (a tuple) negate() tuple {
	return tuple{
		0 - a.x,
		0 - a.y,
		0 - a.z,
		0 - a.w,
	}
}
