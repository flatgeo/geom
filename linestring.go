package geom

import (
	"bytes"
	"strconv"
)

type LineString struct {
	Coordinates []float64
	Extra       int
}

func (line *LineString) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"LineString","coordinates":[`)

	stride := 2 + line.Extra
	for i := 0; i < len(line.Coordinates); i += stride {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)
		for j := 0; j < stride; j += 1 {
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
