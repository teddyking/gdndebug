package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os/exec"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("CLI", func() {
	It("exits with a 0 status code", func() {
		command := exec.Command(pathToGdndebugCLI)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
	})
})
