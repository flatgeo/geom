package geom

import (
	"bytes"
	"strconv"
)

// LineString represents a connected sequence of points.
// Extra dimensions are allowed and must be consistent for all points in the line.
type LineString struct {
	Coordinates []float64
	Extra       int
}

// MarshalJSON returns the GeoJSON representation of the line.
func (line *LineString) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"LineString","coordinates":[`)

	stride := 2 + line.Extra
	for i := 0; i < len(line.Coordinates); i += stride {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)
		for j := 0; j < stride; j++ {
			if j != 0 {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(strconv.FormatFloat(line.Coordinates[i+j], 'g', -1, 64))
		}
		buffer.WriteString(`]`)
	}

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
}
