package integration_test

import (
	"os/exec"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/m-rcd/go-rest-api/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	var (
		session *gexec.Session
		db      database.Database
	)
	BeforeSuite(func() {
		var err error
		var cliBin string
		err = godotenv.Load("./../../.env")
		Expect(err).NotTo(HaveOccurred())
		cliBin, err = gexec.Build("github.com/m-rcd/go-rest-api")
		Expect(err).NotTo(HaveOccurred())
		command := exec.Command(cliBin)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		time.Sleep(2 * time.Second)
		db = database.NewSQL()
		db.Open()
	})

	AfterSuite(func() {
		db.Close()
		session.Terminate().Wait()
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Integration Suite")
}
