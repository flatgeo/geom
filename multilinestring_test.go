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

})
