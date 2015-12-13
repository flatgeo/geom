package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point", func() {

	Describe("json.Marshal", func() {

		It("Encodes points as GeoJSON", func() {
			point := &Point{[]float64{-180, -90}}

			expected, _ := json.Marshal(&geoJSONPoint{
				"Point",
				[]float64{-180, -90},
			})

			Ω(json.Marshal(point)).Should(Equal(expected))
		})

		It("Preserves extra dimensions", func() {
			point := &Point{[]float64{-180, -90, 1.234}}

			expected, _ := json.Marshal(&geoJSONPoint{
				"Point",
				[]float64{-180, -90, 1.234},
			})

			Ω(json.Marshal(point)).Should(Equal(expected))
		})

	})

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON points", func() {
			point := &Point{}

			err := json.Unmarshal([]byte(`{
				"type": "Point",
				"coordinates": [-115, 45]
			}`), point)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(point.Coordinates).Should(Equal([]float64{-115, 45}))
		})

	})

})
