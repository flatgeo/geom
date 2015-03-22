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

	expected := []byte(`{"type":"Point","coordinates":[-180,-90]}`)

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

	expected := []byte(`{"type":"Point","coordinates":[-180,-90,1.234]}`)

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}
