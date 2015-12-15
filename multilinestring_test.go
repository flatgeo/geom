package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiLineString", func() {

	Describe("json.Marshal", func() {

		It("Encodes a MultiLineString as GeoJSON", func() {
			multi := &MultiLineString{[]LineString{
				LineString{Coordinates: []float64{-180, -90, 180, 90}},
				LineString{Coordinates: []float64{-115, 45, 115, -45}},
			}}

			Ω(json.Marshal(multi)).Should(MatchJSON(`{
				"type": "MultiLineString",
				"coordinates": [
					[[-180, -90], [180, 90]],
					[[-115, 45], [115, -45]]
				]
			}`))
		})

	})

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON a MultiLineString", func() {
			multi := &MultiLineString{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiLineString",
				"coordinates": [
					[[-180, -90], [180, 90]],
					[[-115, 45], [115, -45]]
				]
			}`), multi)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(multi.LineStrings).Should(Equal([]LineString{
				LineString{Coordinates: []float64{-180, -90, 180, 90}},
				LineString{Coordinates: []float64{-115, 45, 115, -45}},
			}))
		})

		It("Preserves extra dimensions", func() {
			multi := &MultiLineString{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiLineString",
				"coordinates": [
					[[-180, -90], [180, 90]],
					[[-115, 45, 1.23], [115, -45, 4.56]]
				]
			}`), multi)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(multi.LineStrings).Should(Equal([]LineString{
				LineString{Coordinates: []float64{-180, -90, 180, 90}},
				LineString{
					Coordinates: []float64{-115, 45, 1.23, 115, -45, 4.56},
					Extra:       1,
				},
			}))
		})

		It("Fails for invalid GeoJSON (bad type)", func() {
			multi := &MultiLineString{}

			err := json.Unmarshal([]byte(`{
				"type": "Bogus",
				"coordinates": [
					[[-180, -90], [180, 90]],
					[[-115, 45], [115, -45]]
				]
			}`), multi)

			Ω(err).Should(MatchError(`Unexpected type: "Bogus"`))
		})

		It("Fails for invalid GeoJSON (missing coordinates)", func() {
			multi := &MultiLineString{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiLineString"
			}`), multi)

			Ω(err).Should(MatchError("Expected a coordinate array with one or more values"))
		})

		It("Fails for invalid GeoJSON (inconsistent coordinate dimensions)", func() {
			multi := &MultiLineString{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiLineString",
				"coordinates": [
					[[-180, -90, 1.23], [180, 90]],
					[[-115, 45], [115, -45]]
				]
			}`), multi)

			Ω(err).Should(MatchError("Unexpected coordinates for line 0: Unexpected point length for position 1"))
		})

	})

})
