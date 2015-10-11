package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	binPath string
)

func TestCfExampleBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfExampleBroker Suite")
}

var _ = BeforeSuite(func() {
	var err error
	binPath, err = gexec.Build("github.com/avade/cf-example-broker")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func execBin(args ...string) *gexec.Session {
	cmd := exec.Command(binPath, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	return session
}
