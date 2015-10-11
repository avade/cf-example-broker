package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("CfExampleBroker", func() {

	var (
		session *gexec.Session
	)

	JustBeforeEach(func() {
		args := []string{}

		session = execBin(args...)
	})

	It("exits with a zero exit code", func() {
		Eventually(session).Should(gexec.Exit(0))
	})
})
