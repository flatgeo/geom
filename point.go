package geom

import (
	"bytes"
	"encoding/json"
	"errors"
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

// UnmarshalJSON creates a Point from GeoJSON.
func (point *Point) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONPoint{}
	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "Point" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 2 {
		return errors.New("Expected a coordinates array with two or more values")
	}
	point.Coordinates = geoJSON.Coordinates
	return nil
}
