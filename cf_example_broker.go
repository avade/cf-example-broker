package main

import (
	"errors"
	"net/http"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
)

type exampleServiceBroker struct{}

func (*exampleServiceBroker) Services() []brokerapi.Service {
	return nil
}

func (serviceBroker *exampleServiceBroker) Provision(instanceID string, details brokerapi.ProvisionDetails) error {
	if details.PlanID == serviceBroker.plan().ID {
		return nil
	} else {
		return errors.New("plan_id is not valid")
	}
}

func (*exampleServiceBroker) Deprovision(instanceID string) error {
	return errors.New("Not supported")
	// Deprovision instances here
}

func (*exampleServiceBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (interface{}, error) {
	return nil, errors.New("Not supported")
	// Bind to instances here
	// Return credentials which will be marshalled to JSON
}

func (*exampleServiceBroker) Unbind(instanceID, bindingID string) error {
	return errors.New("Not supported")
	// Unbind from instances here
}

func main() {
	serviceBroker := &exampleServiceBroker{}
	logger := lager.NewLogger("cf-example-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)
	http.Handle("/", brokerAPI)
	http.ListenAndServe(":3000", nil)
}

func (serviceBroker *exampleServiceBroker) plan() *brokerapi.ServicePlan {
	return &brokerapi.ServicePlan{
		ID:          "cheap-id",
		Name:        "cheap",
		Description: "This plan provides a...",
		Metadata: brokerapi.ServicePlanMetadata{
			Bullets: []string{
				"Example CF service",
			},
			DisplayName: "Cheap-Plan",
		},
	}
}
