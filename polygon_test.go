package geom

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPolygonJSON(t *testing.T) {
	poly := &Polygon{
		Coordinates: []float64{-180, -90, 180, -90, 180, 90, -180, 90, -180, -90},
	}

	got, err := json.Marshal(poly)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _ := json.Marshal(struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	}{
		"Polygon",
		[][][]float64{
			{{-180, -90}, {180, -90}, {180, 90}, {-180, 90}, {-180, -90}},
		},
	})

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}

func TestPolygonJSONHoles(t *testing.T) {
	poly := &Polygon{
		Coordinates: []float64{
			-180, -90, 180, -90, 180, 90, -180, 90, -180, -90,
			-100, -45, -100, 45, -50, 45, -50, -45, -100, -45,
			100, -45, 100, 45, 50, 45, 50, -45, 100, -45,
		},
		RingStarts: []int{10, 20},
	}

	got, err := json.Marshal(poly)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _ := json.Marshal(struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	}{
		"Polygon",
		[][][]float64{
			{{-180, -90}, {180, -90}, {180, 90}, {-180, 90}, {-180, -90}},
			{{-100, -45}, {-100, 45}, {-50, 45}, {-50, -45}, {-100, -45}},
			{{100, -45}, {100, 45}, {50, 45}, {50, -45}, {100, -45}},
		},
	})

	if !bytes.Equal(got, expected) {
		t.Errorf("bad json: got %s but expected %s", got, expected)
	}
}
