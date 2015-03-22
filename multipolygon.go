package geom

type MultiPolygon struct {
	Coordinates []float64
	PolyStarts  []int
	RingStarts  []int
	Extra       int
}
