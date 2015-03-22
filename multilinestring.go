package geom

type MultiLineString struct {
	Coordinates []float64
	LineStarts  []int
	Extra       int
}
