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

			Ω(json.Marshal(line)).Should(MatchJSON(`{
				"type": "LineString",
				"coordinates": [[-180, -90], [180, 90]]
			}`))
		})

		It("Preserves extra dimensions", func() {
			line := &LineString{
				Coordinates: []float64{-180, -90, 1.23, 180, 90, 4.56},
				Extra:       1,
			}

			Ω(json.Marshal(line)).Should(MatchJSON(`{
				"type": "LineString",
				"coordinates": [[-180, -90, 1.23], [180, 90, 4.56]]
			}`))
		})

	})

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON linestrings", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "LineString",
				"coordinates": [[-115, 45], [115, 45]]
			}`), line)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(line.Coordinates).Should(Equal([]float64{-115, 45, 115, 45}))
			Ω(line.Extra).Should(Equal(0))
		})

		It("Preserves extra dimensions", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "LineString",
				"coordinates": [[-115, 45, 1.234], [115, 45, 5.678]]
			}`), line)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(line.Coordinates).Should(Equal([]float64{-115, 45, 1.234, 115, 45, 5.678}))
			Ω(line.Extra).Should(Equal(1))
		})

		It("Fails for invalid GeoJSON (bad type)", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "Bogus",
				"coordinates": [[-115, 45], [115, 45]]
			}`), line)

			Ω(err).Should(MatchError(`Unexpected type: "Bogus"`))
		})

		It("Fails for invalid GeoJSON (missing coordinates)", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "LineString"
			}`), line)

			Ω(err).Should(MatchError("Expected a coordinates array with two or more values"))
		})

		It("Fails for invalid GeoJSON (not enough coordinates)", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "LineString",
				"coordinates": [[1, 2]]
			}`), line)

			Ω(err).Should(MatchError("Expected a coordinates array with two or more values"))
		})

		It("Fails for invalid GeoJSON (inconsistent coordinate dimensions)", func() {
			line := &LineString{}

			err := json.Unmarshal([]byte(`{
				"type": "LineString",
				"coordinates": [[1, 2], [3, 4, 5]]
			}`), line)

			Ω(err).Should(MatchError("Unexpected point length for position 1"))
		})

	})

})
