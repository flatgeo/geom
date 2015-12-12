package geom

// MultiPolygon represents a collection of polygons.
type MultiPolygon struct {
	Coordinates []float64
	PolyStarts  []int
	RingStarts  []int
	Extra       int
}
