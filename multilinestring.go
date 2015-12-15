package geom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

// UnmarshalJSON creates a MultiLineString from GeoJSON.
func (multi *MultiLineString) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONMultiLineString{}
	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "MultiLineString" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 1 {
		return errors.New("Expected a coordinate array with one or more values")
	}

	lines := make([]LineString, len(geoJSON.Coordinates))
	for i, coords := range geoJSON.Coordinates {
		line := LineString{}
		if err := line.setCoordinates(coords); err != nil {
			return fmt.Errorf("Unexpected coordinates for line %d: %s", i, err.Error())
		}
		lines[i] = line
	}
	multi.LineStrings = lines
	return nil
}
