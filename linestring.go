package geom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// LineString represents a connected sequence of points.
// Extra dimensions are allowed and must be consistent for all points in the line.
type LineString struct {
	Coordinates []float64
	Extra       int
}

type geoJSONLineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

// MarshalJSON returns the GeoJSON encoding of a LineString.
func (line *LineString) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"LineString","coordinates":[`)

	dimensions := 2 + line.Extra
	for i := 0; i < len(line.Coordinates); i += dimensions {
		if i != 0 {
			buffer.WriteString(`,`)
		}
		buffer.WriteString(`[`)
		for j := 0; j < dimensions; j++ {
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

// UnmarshalJSON creates a LineString from GeoJSON.
func (line *LineString) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONLineString{}
	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "LineString" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 2 {
		return errors.New("Expected a coordinates array with two or more values")
	}

	return line.setCoordinates(geoJSON.Coordinates)
}

func (line *LineString) setCoordinates(coords [][]float64) error {
	first := coords[0]
	dimensions := len(first)
	if dimensions < 2 {
		return fmt.Errorf("Unexpected length %d for first point", dimensions)
	}

	coordinates := make([]float64, len(coords)*dimensions)
	for i, coord := range coords {
		if len(coord) != dimensions {
			return fmt.Errorf("Unexpected point length for position %d", i)
		}
		for j := 0; j < dimensions; j++ {
			coordinates[i*dimensions+j] = coord[j]
		}
	}

	line.Coordinates = coordinates
	line.Extra = dimensions - 2
	return nil
}
