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

			Ω(json.Marshal(poly)).Should(MatchJSON(`{
				"type": "Polygon",
				"coordinates": [
					[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]
				]
			}`))
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

			Ω(json.Marshal(poly)).Should(MatchJSON(`{
				"type": "Polygon",
				"coordinates": [
					[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]],
					[[-100, -45], [-100, 45], [-50, 45], [-50, -45], [-100, -45]],
					[[100, -45], [100, 45], [50, 45], [50, -45], [100, -45]]
				]
			}`))
		})

	})

})
