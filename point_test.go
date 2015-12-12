package geom

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPointJSON(t *testing.T) {
	point := &Point{[]float64{-180, -90}}

	got, err := json.Marshal(point)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _ := json.Marshal(&geoJSONPoint{
		"Point",
		[]float64{-180, -90},
	})

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}

func TestPointJSON3D(t *testing.T) {
	point := &Point{[]float64{-180, -90, 1.234}}

	got, err := json.Marshal(point)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _ := json.Marshal(&geoJSONPoint{
		"Point",
		[]float64{-180, -90, 1.234},
	})

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}

func TestPointUnmarshal(t *testing.T) {
	point := &Point{}

	err := json.Unmarshal([]byte(`{
		"type": "Point",
		"coordinates": [-115, 45]
	}`), point)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(point.Coordinates) != 2 {
		t.Fatalf("expected coordinates of length 2, got %f", len(point.Coordinates))
	}

	if point.Coordinates[0] != -115 {
		t.Fatalf("expected coordinates[0] to be -115, got %f", point.Coordinates[0])
	}

	if point.Coordinates[1] != 45 {
		t.Fatalf("expected coordinates[0] to be 45, got %f", point.Coordinates[0])
	}

}
