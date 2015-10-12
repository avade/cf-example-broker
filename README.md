#cf-example-broker

An example service broker for Cloud Foundry. 

Use the projects tags to jump through milestones of the development. 

## Running the Broker

### Configuration

First update the configuration file for the credentials to hand out on bind.

```
vim assets/config.yml

```

```
service_broker:
 max_instances: 5
 username: user
 password: password
service_credentials:
 username: secret
 password: service
```

### Build & Run

```
go build
./cf-example-broker config.yml
```

The broker is now running on `http://localhost:3000`

## Running the tests


- Setup [ginkgo](https://onsi.github.io/ginkgo/#getting-ginkgo)
- `ginkgo -r`

## Libaries

This section links to OSS libaries that have been used in the project.

- [Ginkgo](https://onsi.github.io/ginkgo) 
- [Gomega](https://onsi.github.io/gomega)
- [Broker API](https://github.com/pivotal-cf/brokerapi)

## Links

- [Broker API - CF Docs](https://docs.cloudfoundry.org/services/api.html)
- [Managing Service Brokers - CF Docs](https://docs.cloudfoundry.org/services/managing-service-brokers.html)
- [CF Redis Broker](https://github.com/pivotal-cf/cf-redis-broker)
- [Pivotal Golang GitHub](https://github.com/pivotal-golang)
- [GoKit - A distributed programming toolkit](https://github.com/go-kit/kit)
