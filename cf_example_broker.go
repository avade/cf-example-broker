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

func (*exampleServiceBroker) Provision(instanceID string, details brokerapi.ProvisionDetails) error {
	return errors.New("Not supported")
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
