package geom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Polygon represents a surface enclosed by an exterior ring (with optional holes).
type Polygon struct {
	Coordinates []float64
	RingStarts  []int
	Extra       int
}

type geoJSONPolygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// MarshalJSON returns the GeoJSON encoding for a Polygon.
func (poly *Polygon) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"Polygon","coordinates":`)
	poly.writeCoordinates(&buffer)
	buffer.WriteString(`}`)
	return buffer.Bytes(), nil
}

func (poly *Polygon) writeCoordinates(buffer *bytes.Buffer) {
	buffer.WriteString(`[`)
	dimensions := 2 + poly.Extra
	start := 0
	for ring := 0; ring <= len(poly.RingStarts); ring++ {
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
		for i := start; i < end; i += dimensions {
			if i != start {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(`[`)
			for j := 0; j < dimensions; j++ {
				if j != 0 {
					buffer.WriteString(`,`)
				}
				buffer.WriteString(strconv.FormatFloat(poly.Coordinates[i+j], 'g', -1, 64))
			}
			buffer.WriteString(`]`)
		}
		buffer.WriteString(`]`)
		start = end
	}
	buffer.WriteString(`]`)
}

// UnmarshalJSON creates a Polygon from GeoJSON
func (poly *Polygon) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONPolygon{}

	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "Polygon" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 1 {
		return errors.New("Expected a coordinates array with one or more rings")
	}

	return poly.setCoordinates(geoJSON.Coordinates)
}

func (poly *Polygon) setCoordinates(coords [][][]float64) error {
	var coordinates []float64
	var dimensions int
	var ringStarts []int

	for r, ring := range coords {
		if r != 0 {
			ringStarts = append(ringStarts, len(coordinates))
		}
		for i, coord := range ring {
			if r == 0 && i == 0 {
				dimensions = len(coord)
				if dimensions < 2 {
					return fmt.Errorf(`Unexpected length %d for first point`, dimensions)
				}
			} else {
				if len(coord) != dimensions {
					return fmt.Errorf(`Unexpected length %d for point %d in ring %d`, len(coord), i, r)
				}
			}
			for j := 0; j < dimensions; j++ {
				coordinates = append(coordinates, coord[j])
			}
		}
	}

	poly.Coordinates = coordinates
	poly.Extra = dimensions - 2
	poly.RingStarts = ringStarts
	return nil
}
