package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LineString", func() {

	Describe("json.Marshal", func() {

		It("Encodes lines as GeoJSON", func() {
			line := &LineString{
				Coordinates: []float64{-180, -90, 180, 90},
			}

			expected, _ := json.Marshal(struct {
				Type        string      `json:"type"`
				Coordinates [][]float64 `json:"coordinates"`
			}{
				"LineString",
				[][]float64{{-180, -90}, {180, 90}},
			})

			Ω(json.Marshal(line)).Should(Equal(expected))
		})

		It("Preserves extra dimensions", func() {
			line := &LineString{
				Coordinates: []float64{-180, -90, 1.23, 180, 90, 4.56},
				Extra:       1,
			}

			expected, _ := json.Marshal(struct {
				Type        string      `json:"type"`
				Coordinates [][]float64 `json:"coordinates"`
			}{
				"LineString",
				[][]float64{{-180, -90, 1.23}, {180, 90, 4.56}},
			})

			Ω(json.Marshal(line)).Should(Equal(expected))
		})

	})

})
