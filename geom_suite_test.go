package geom_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGeom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geom Suite")
}
