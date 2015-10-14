package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	provisionRequest := func(provisionDetails brokerapi.ProvisionDetails) (*http.Response, error) {
		client := &http.Client{}
		path := fmt.Sprintf("http://localhost:3000/v2/service_instances/%s", provisionDetails.ID)

		buffer := &bytes.Buffer{}
		json.NewEncoder(buffer).Encode(provisionDetails)

		request, err := http.NewRequest("PUT", path, buffer)
		Expect(err).NotTo(HaveOccurred())

		request.Header.Add("Content-Type", "application/json")
		request.SetBasicAuth("username", "password")

		return client.Do(request)
	}

	Describe(".Provision", func() {

		var (
			provisionDetails brokerapi.ProvisionDetails
			response         *http.Response
			err              error
		)

		JustBeforeEach(func() {
			response, err = provisionRequest(provisionDetails)
		})

		Context("when the correct plan is provided", func() {

			BeforeEach(func() {
				provisionDetails = brokerapi.ProvisionDetails{
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
				provisionDetails = brokerapi.ProvisionDetails{
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

	Describe(".Bind", func() {

		var (
			provisionDetails brokerapi.ProvisionDetails
			bindDetails      brokerapi.BindDetails
			response         *http.Response
			err              error
		)

		JustBeforeEach(func() {
			_, err := provisionRequest(provisionDetails)
			Expect(err).NotTo(HaveOccurred())

			//create the binding
			client := &http.Client{}
			path := fmt.Sprintf("http://localhost:3000/v2/service_instances/%s/service_bindings/1", bindDetails.ServiceID)

			buffer := &bytes.Buffer{}
			json.NewEncoder(buffer).Encode(bindDetails)

			request, err := http.NewRequest("PUT", path, buffer)
			Expect(err).NotTo(HaveOccurred())

			request.Header.Add("Content-Type", "application/json")
			request.SetBasicAuth("username", "password")

			response, err = client.Do(request)
		})

		Context("when the service instance exists", func() {

			Context("when the binding is new", func() {
				BeforeEach(func() {
					provisionDetails = brokerapi.ProvisionDetails{
						ID:               "service-id",
						PlanID:           "cheap-id",
						OrganizationGUID: "organization-guid",
						SpaceGUID:        "space-guid",
					}
					bindDetails = brokerapi.BindDetails{
						ServiceID: "service-id",
						PlanID:    "cheap-id",
						AppGUID:   "123",
					}
				})

				It("does not return an error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("responds with a 201", func() {
					Expect(response.StatusCode).To(Equal(http.StatusCreated))
				})

				It("responds with the credentials", func() {
					body, err := ioutil.ReadAll(response.Body)
					Expect(err).ToNot(HaveOccurred())
					var jsonData struct {
						Credentials struct {
							Username string
							Password string
						}
					}
					err = json.Unmarshal(body, &jsonData)
					Expect(err).ToNot(HaveOccurred())

					Expect(jsonData.Credentials.Username).ToNot(BeEmpty())
					Expect(jsonData.Credentials.Password).ToNot(BeEmpty())
				})
			})
		})

		Context("when the service doesn't exist", func() {
			BeforeEach(func() {
				provisionDetails = brokerapi.ProvisionDetails{
					ID:               "service-id",
					PlanID:           "cheap-id",
					OrganizationGUID: "organization-guid",
					SpaceGUID:        "space-guid",
				}
				bindDetails = brokerapi.BindDetails{
					ServiceID: "wrong-id",
					PlanID:    "cheap-id",
					AppGUID:   "123",
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
