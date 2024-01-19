package tuple

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
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
