package geom

import (
	"bytes"
	"strconv"
)

// MultiLineString represents a collection of lines.
type MultiLineString struct {
	LineStrings []LineString
}

type geoJSONMultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// MarshalJSON returns the GeoJSON encoding of a MultiLineString.
func (multi *MultiLineString) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer

	buffer.WriteString(`{"type":"MultiLineString","coordinates":[`)

	for i, line := range multi.LineStrings {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)

		dimensions := line.Extra + 2
		for j := 0; j < len(line.Coordinates); j += dimensions {
			if j != 0 {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(`[`)
			for k := 0; k < dimensions; k++ {
				if k != 0 {
					buffer.WriteString(`,`)
				}
				buffer.WriteString(strconv.FormatFloat(line.Coordinates[j+k], 'g', -1, 64))
			}
			buffer.WriteString(`]`)
		}

		buffer.WriteString(`]`)
	}

	buffer.WriteString(`]}`)

	return buffer.Bytes(), nil
}
