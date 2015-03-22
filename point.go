package geom

import (
	"bytes"
	"strconv"
)

type Point struct {
	Coordinates []float64
}

func (point *Point) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer

	buffer.WriteString(`{"type":"Point","coordinates":[`)
	for i, v := range point.Coordinates {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(strconv.FormatFloat(v, 'g', -1, 64))
	}

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
}
