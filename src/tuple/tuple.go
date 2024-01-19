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
