package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var pathToGdndebugCLI string

var _ = BeforeSuite(func() {
	var err error
	pathToGdndebugCLI, err = gexec.Build("github.com/teddyking/gdndebug/cmd/gdndebug")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}
