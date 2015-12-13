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

})
