package geom

import (
	"bytes"
	"strconv"
)

// Polygon represents a surface enclosed by an exterior ring (with optional holes).
type Polygon struct {
	Coordinates []float64
	RingStarts  []int
	Extra       int
}

// MarshalJSON returns the GeoJSON encoding for the polygon.
func (poly *Polygon) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"Polygon","coordinates":[`)

	stride := 2 + poly.Extra
	start := 0
	for ring := 0; ring <= len(poly.RingStarts); ring++ {
		if ring != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)
		var end int
		if ring < len(poly.RingStarts) {
			end = poly.RingStarts[ring]
		} else {
			end = len(poly.Coordinates)
		}
		for i := start; i < end; i += stride {
			if i != start {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(`[`)
			for j := 0; j < stride; j++ {
				if j != 0 {
					buffer.WriteString(`,`)
				}
				buffer.WriteString(strconv.FormatFloat(poly.Coordinates[i+j], 'g', -1, 64))
			}
			buffer.WriteString(`]`)
		}
		buffer.WriteString(`]`)
		start = end
	}

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
}
