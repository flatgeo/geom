package geom

import (
	"bytes"
	"strconv"
)

type Polygon struct {
	Coordinates []float64
	RingStarts  []int
	Extra       int
}

func (poly *Polygon) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"Polygon","coordinates":[`)

	stride := 2 + poly.Extra
	start := 0
	for ring := 0; ring <= len(poly.RingStarts); ring += 1 {
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
			if i != 0 {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(`[`)
			for j := 0; j < stride; j += 1 {
				if j != 0 {
					buffer.WriteString(`,`)
				}
				buffer.WriteString(strconv.FormatFloat(poly.Coordinates[i+j], 'g', -1, 64))
			}
			buffer.WriteString(`]`)
		}
		buffer.WriteString(`]`)
		start = end + 1
	}

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
}
