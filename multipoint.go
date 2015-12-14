package geom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// MultiPoint represents a collection of points.
type MultiPoint struct {
	Points []Point
}

type geoJSONMultiPoint struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

// MarshalJSON returns the GeoJSON encoding of a multi-point.
func (multi *MultiPoint) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer

	buffer.WriteString(`{"type":"MultiPoint","coordinates":[`)
	for i, point := range multi.Points {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)
		for j, value := range point.Coordinates {
			if j != 0 {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(strconv.FormatFloat(value, 'g', -1, 64))
		}
		buffer.WriteString(`]`)
	}

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON creates a MultiPoint from GeoJSON.
func (multi *MultiPoint) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONMultiPoint{}
	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "MultiPoint" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 1 {
		return errors.New("Expected a coordinates array with one or more values")
	}
	var points []Point
	for _, coord := range geoJSON.Coordinates {
		points = append(points, Point{coord})
	}
	multi.Points = points
	return nil
}
