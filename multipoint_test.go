package geom

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiPoint", func() {

	Describe("json.Marshal", func() {

		It("Encodes MultiPoint as GeoJSON", func() {
			multi := &MultiPoint{[]Point{
				Point{[]float64{-180, -90}},
				Point{[]float64{180, 90}},
			}}

			Ω(json.Marshal(multi)).Should(MatchJSON(`{
				"type": "MultiPoint",
				"coordinates": [[-180, -90], [180, 90]]
			}`))
		})

		It("Preserves extra dimensions", func() {
			multi := &MultiPoint{[]Point{
				Point{[]float64{-180, -90, 1.234}},
				Point{[]float64{180, 90, 5.678}},
			}}

			Ω(json.Marshal(multi)).Should(MatchJSON(`{
				"type": "MultiPoint",
				"coordinates": [[-180, -90, 1.234], [180, 90, 5.678]]
			}`))
		})

	})

	Describe("json.Unmarshal", func() {

		It("Decodes GeoJSON MultiPoint", func() {
			multi := &MultiPoint{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiPoint",
				"coordinates": [[-115, 45], [115, -45]]
			}`), multi)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(multi.Points).Should(Equal([]Point{
				Point{[]float64{-115, 45}},
				Point{[]float64{115, -45}},
			}))
		})

		It("Preserves extra dimensions", func() {
			multi := &MultiPoint{}

			err := json.Unmarshal([]byte(`{
				"type": "MultiPoint",
				"coordinates": [[-115, 45, 1.234], [115, -45, 5.678]]
			}`), multi)

			Ω(err).ShouldNot(HaveOccurred())

			Ω(multi.Points).Should(Equal([]Point{
				Point{[]float64{-115, 45, 1.234}},
				Point{[]float64{115, -45, 5.678}},
			}))
		})

		It("Fails for invalid GeoJSON (bad type)", func() {
			multi := &MultiPoint{}

			err := json.Unmarshal([]byte(`{
				"type": "Bogus",
				"coordinates": [[-115, 45], [115, -45]]
			}`), multi)

			Ω(err).Should(MatchError(`Unexpected type: "Bogus"`))
		})

	})

})
