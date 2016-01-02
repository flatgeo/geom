package geom

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MultiPolygon represents a collection of polygons.
type MultiPolygon struct {
	Polygons []Polygon
}

type geoJSONMultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
}

// UnmarshalJSON creates a MultiPolygon from GeoJSON
func (multi *MultiPolygon) UnmarshalJSON(data []byte) error {
	geoJSON := &geoJSONMultiPolygon{}
	if err := json.Unmarshal(data, geoJSON); err != nil {
		return err
	}
	if geoJSON.Type != "MultiPolygon" {
		return fmt.Errorf(`Unexpected type: "%s"`, geoJSON.Type)
	}
	if len(geoJSON.Coordinates) < 1 {
		return errors.New("Expected a coordinate array with one or more values")
	}

	polygons := make([]Polygon, len(geoJSON.Coordinates))
	for i, coords := range geoJSON.Coordinates {
		poly := Polygon{}
		if err := poly.setCoordinates(coords); err != nil {
			return fmt.Errorf("Unexpected coordinates for polygon %d: %s", i, err.Error())
		}
		polygons[i] = poly
	}
	multi.Polygons = polygons
	return nil
}
