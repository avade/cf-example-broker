package database

import (
	"github.com/satori/go.uuid"
)

type Creator interface {
	CreateDb() (error, string)
	CreateUser(string) (error, string, string)
}

type service struct {
	username string
	password string
	hostname string
	port     uint64
}

func NewCreator(username, password, hostname string, port uint64) Creator {
	return &service{
		username: username,
		password: password,
		hostname: hostname,
		port:     port,
	}
}

func (service *service) CreateDb() (error, string) {
	return nil, uuid.NewV4().String()
}

func (service *service) CreateUser(databaseName string) (error, string, string) {
	return nil, uuid.NewV4().String(), uuid.NewV4().String()
}
