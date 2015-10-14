#cf-example-broker

An example service broker for Cloud Foundry. The purpose is to showcase a range of OSS libraries that help speed up Broker development.

This broker does not connect to an actual database, but eventually it will hand out bindings and conform to the [Broker API](https://docs.cloudfoundry.org/services/api.html).

Use the projects tags to jump through milestones of the development.

## Running the Broker

### Configuration (not implemented)

First update the configuration file. 

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
 host: host
 port: 1234
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

## Contributing

Please send any contributions via a pull request or raise an [issue].

## Libaries

This section links to OSS libaries that have been used in the project.

- [Ginkgo](https://onsi.github.io/ginkgo) 
- [Gomega](https://onsi.github.io/gomega)
- [Broker API](https://github.com/pivotal-cf/brokerapi)
- [Counterfeiter - Fake generator](https://github.com/maxbrunsfeld/counterfeiter)

## Links

- [Broker API - CF Docs](https://docs.cloudfoundry.org/services/api.html)
- [Managing Service Brokers - CF Docs](https://docs.cloudfoundry.org/services/managing-service-brokers.html)
- [CF Redis Broker](https://github.com/pivotal-cf/cf-redis-broker)
- [Pivotal Golang GitHub](https://github.com/pivotal-golang)
- [GoKit - A distributed programming toolkit](https://github.com/go-kit/kit)
