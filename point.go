package geom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Point represents coordinates of a single position.
type Point struct {
	Coordinates []float64
}

type geoJSONPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON returns the GeoJSON encoding of a point.
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

// UnmarshalJSON creates a point from GeoJSON.
func (point *Point) UnmarshalJSON(data []byte) error {
	geoJSON := geoJSONPoint{}
	if err := json.Unmarshal(data, &geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "Point" {
		return fmt.Errorf(`Expected "type": "Point", got: %s`, geoJSON.Type)
	}
	point.Coordinates = geoJSON.Coordinates
	return nil
}
