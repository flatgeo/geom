package geom

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPolygonJSON(t *testing.T) {
	poly := &Polygon{
		Coordinates: []float64{-180, -90, 180, -90, 180, 90, -180, 90},
	}

	got, err := json.Marshal(poly)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []byte(`{"type":"Polygon","coordinates":[[[-180,-90],[180,-90],[180,90],[-180,90]]]}`)

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}
