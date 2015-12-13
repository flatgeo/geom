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

// MarshalJSON returns the GeoJSON encoding for the polygon.
func (poly *Polygon) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"type":"Polygon","coordinates":[`)

	stride := 2 + poly.Extra
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
		for i := start; i < end; i += stride {
			if i != start {
				buffer.WriteString(`,`)
			}
			buffer.WriteString(`[`)
			for j := 0; j < stride; j++ {
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

	buffer.WriteString(`]}`)
	return buffer.Bytes(), nil
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

	var coordinates []float64
	var dimensions int
	var ringStarts []int

	for r := 0; r < len(geoJSON.Coordinates); r++ {
		ring := geoJSON.Coordinates[r]
		if r != 0 {
			ringStarts = append(ringStarts, len(coordinates))
		}
		for i := 0; i < len(ring); i++ {
			coord := ring[i]
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
