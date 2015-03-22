package geom

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestLineStringJSON(t *testing.T) {
	line := &LineString{
		Coordinates: []float64{-180, -90, 180, 90},
	}

	got, err := json.Marshal(line)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []byte(`{"type":"LineString","coordinates":[[-180,-90],[180,90]]}`)

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}

func TestLineStringJSON3D(t *testing.T) {
	line := &LineString{
		Coordinates: []float64{-180, -90, 1.23, 180, 90, 4.56},
		Extra:       1,
	}

	got, err := json.Marshal(line)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []byte(`{"type":"LineString","coordinates":[[-180,-90,1.23],[180,90,4.56]]}`)

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}
