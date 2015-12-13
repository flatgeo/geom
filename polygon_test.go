package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Polygon", func() {

	Describe("json.Marshal", func() {

		It("Encodes polygons as GeoJSON", func() {

			poly := &Polygon{
				Coordinates: []float64{-180, -90, 180, -90, 180, 90, -180, 90, -180, -90},
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

			Ω(json.Marshal(poly)).Should(Equal(expected))
		})

		It("Works with interior rings", func() {

			poly := &Polygon{
				Coordinates: []float64{
					-180, -90, 180, -90, 180, 90, -180, 90, -180, -90,
					-100, -45, -100, 45, -50, 45, -50, -45, -100, -45,
					100, -45, 100, 45, 50, 45, 50, -45, 100, -45,
				},
				RingStarts: []int{10, 20},
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

			Ω(json.Marshal(poly)).Should(Equal(expected))
		})

	})

})
