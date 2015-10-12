package database_test

import (
	"github.com/avade/cf-example-broker/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database Service", func() {

	var (
		creator database.Creator
	)

	BeforeEach(func() {
		creator = database.NewCreator("username", "password", "hostname", 1234)
	})

	Context("Creating a database", func() {
		It("Doesn't return an error", func() {
			err, _ := creator.CreateDb()
			Expect(err).ToNot(HaveOccurred())
		})

		It("Returns the database name", func() {
			_, dbName := creator.CreateDb()
			Expect(dbName).ToNot(BeNil())
		})
	})

	Context("Creating a user for a given database", func() {
		It("Doesn't return an error", func() {
			err, _, _ := creator.CreateUser("database")
			Expect(err).ToNot(HaveOccurred())
		})

		It("Returns the username & password", func() {
			_, username, password := creator.CreateUser("database")
			Expect(username).ToNot(BeNil())
			Expect(password).ToNot(BeNil())
		})
	})

})
