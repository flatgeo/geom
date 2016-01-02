package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiPolygon", func() {

	Describe("json.Marshal", func() {

		It("Encodes MultiPolygons as GeoJSON", func() {
			multi := &MultiPolygon{[]Polygon{
				Polygon{
					Coordinates: []float64{-180, -90, 180, -90, 180, 90, -180, 90, -180, -90},
				},
				Polygon{
					Coordinates: []float64{-110, -45, 110, -45, 110, 45, -110, 45, -110, -45},
				},
			}}

			Ω(json.Marshal(multi)).Should(MatchJSON(`{
				"type": "MultiPolygon",
				"coordinates": [
					[[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]],
					[[[-110, -45], [110, -45], [110, 45], [-110, 45], [-110, -45]]]
				]
			}`))
		})

	})

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON a MultiPolygon", func() {
			multi := &MultiPolygon{}

			err := json.Unmarshal([]byte(`{
	 			"type": "MultiPolygon",
	 			"coordinates": [
	 				[[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]],
	 				[[[-110, -45], [110, -45], [110, 45], [-110, 45], [-110, -45]]]
	 			]
			}`), multi)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(multi.Polygons).Should(Equal([]Polygon{
				Polygon{
					Coordinates: []float64{-180, -90, 180, -90, 180, 90, -180, 90, -180, -90},
				},
				Polygon{
					Coordinates: []float64{-110, -45, 110, -45, 110, 45, -110, 45, -110, -45},
				},
			}))
		})

		It("Fails with invalid GeoJSON (bad type)", func() {
			multi := &MultiPolygon{}

			err := json.Unmarshal([]byte(`{
	 			"type": "Bogus",
	 			"coordinates": [
	 				[[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]],
	 				[[[-110, -45], [110, -45], [110, 45], [-110, 45], [-110, -45]]]
	 			]
			}`), multi)

			Ω(err).Should(MatchError(`Unexpected type: "Bogus"`))
		})

		It("Fails with invalid GeoJSON (mixed dimension)", func() {
			multi := &MultiPolygon{}

			err := json.Unmarshal([]byte(`{
	 			"type": "MultiPolygon",
	 			"coordinates": [
	 				[[[-180, -90, 123], [180, -90], [180, 90], [-180, 90], [-180, -90]]],
	 				[[[-110, -45], [110, -45], [110, 45], [-110, 45], [-110, -45]]]
	 			]
			}`), multi)

			Ω(err).Should(MatchError("Unexpected coordinates for polygon 0: Unexpected length 2 for point 1 in ring 0"))
		})

	})

})
