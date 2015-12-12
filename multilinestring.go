package geom

// MultiLineString represents a collection of lines.
type MultiLineString struct {
	Coordinates []float64
	LineStarts  []int
	Extra       int
}
