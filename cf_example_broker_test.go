package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/brokerapi"
)

var _ = Describe("CfExampleBroker", func() {

	var (
		session *gexec.Session
	)

	BeforeEach(func() {
		args := []string{}
		session = execBin(args...)
		time.Sleep(1 * time.Second)
	})

	AfterEach(func() {
		session.Kill()
	})

	Describe(".Provision", func() {

		var (
			details  brokerapi.ProvisionDetails
			response *http.Response
			err      error
		)

		JustBeforeEach(func() {
			client := &http.Client{}
			path := "http://localhost:3000/v2/service_instances/123"

			buffer := &bytes.Buffer{}
			json.NewEncoder(buffer).Encode(details)
			var request *http.Request
			request, err = http.NewRequest("PUT", path, buffer)
			Expect(err).NotTo(HaveOccurred())

			request.Header.Add("Content-Type", "application/json")
			request.SetBasicAuth("username", "password")

			response, err = client.Do(request)
		})

		Context("when the correct plan is provided", func() {

			BeforeEach(func() {
				details = brokerapi.ProvisionDetails{
					ID:               "service-id",
					PlanID:           "cheap-id",
					OrganizationGUID: "organization-guid",
					SpaceGUID:        "space-guid",
				}
			})

			It("does not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("responds with a 201", func() {
				Expect(response.StatusCode).To(Equal(http.StatusCreated))
			})
		})

		Context("when the plan invalid", func() {
			BeforeEach(func() {
				details = brokerapi.ProvisionDetails{
					ID:               "service-id",
					PlanID:           "expensive-id",
					OrganizationGUID: "organization-guid",
					SpaceGUID:        "space-guid",
				}
			})

			It("does not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			//TOOD: brokerapi does not support this error, should 400 BadRequest
			It("responds with a 500", func() {
				Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
