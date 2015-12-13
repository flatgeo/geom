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

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON polygons", func() {
			poly := &Polygon{}

			err := json.Unmarshal([]byte(`{
				"type": "Polygon",
				"coordinates": [
					[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]
				]
			}`), poly)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(poly.Extra).Should(Equal(0))
			Ω(poly.Coordinates).Should(Equal([]float64{-180, -90, 180, -90, 180, 90, -180, 90, -180, -90}))
		})

		It("Works with interior rings", func() {
			poly := &Polygon{}

			err := json.Unmarshal([]byte(`{
				"type": "Polygon",
				"coordinates": [
					[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]],
					[[-100, -45], [-100, 45], [-50, 45], [-50, -45], [-100, -45]],
					[[100, -45], [100, 45], [50, 45], [50, -45], [100, -45]]
				]
			}`), poly)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(poly.Extra).Should(Equal(0))
			Ω(poly.Coordinates).Should(Equal([]float64{
				-180, -90, 180, -90, 180, 90, -180, 90, -180, -90,
				-100, -45, -100, 45, -50, 45, -50, -45, -100, -45,
				100, -45, 100, 45, 50, 45, 50, -45, 100, -45,
			}))
			Ω(poly.RingStarts).Should(Equal([]int{10, 20}))
		})

	})

	It("Preserves extra dimensions", func() {
		poly := &Polygon{}

		err := json.Unmarshal([]byte(`{
			"type": "Polygon",
			"coordinates": [
				[[-180, -90, 1], [180, -90, 2], [180, 90, 3], [-180, 90, 4], [-180, -90, 1]]
			]
		}`), poly)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(poly.Extra).Should(Equal(1))
		Ω(poly.Coordinates).Should(Equal([]float64{-180, -90, 1, 180, -90, 2, 180, 90, 3, -180, 90, 4, -180, -90, 1}))
	})

	It("Fails with invalid GeoJSON (bad type)", func() {
		poly := &Polygon{}

		err := json.Unmarshal([]byte(`{
			"type": "Bogus",
			"coordinates": [
				[[-180, -90], [180, -90], [180, 90], [-180, 90], [-180, -90]]
			]
		}`), poly)

		Ω(err).Should(MatchError(`Unexpected type: "Bogus"`))
	})

	It("Fails with invalid GeoJSON (mixed dimension)", func() {
		poly := &Polygon{}

		err := json.Unmarshal([]byte(`{
			"type": "Polygon",
			"coordinates": [
				[[-180, -90], [180, -90, 123], [180, 90], [-180, 90], [-180, -90]]
			]
		}`), poly)

		Ω(err).Should(MatchError("Unexpected length 3 for point 1 in ring 0"))
	})

})
