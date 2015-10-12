package database

import (
	"errors"

	"github.com/satori/go.uuid"
)

type Creator interface {
	CreateDb() (error, string)
	CreateUser(string) (error, string, string)
}

type service struct {
	username  string
	password  string
	hostname  string
	port      uint64
	databases []string
}

func NewCreator(username, password, hostname string, port uint64) Creator {
	return &service{
		username: username,
		password: password,
		hostname: hostname,
		port:     port,
	}
}

func (s *service) CreateDb() (error, string) {
	dbName := uuid.NewV4().String()
	s.databases = append(s.databases, dbName)
	return nil, dbName
}

func (s *service) CreateUser(databaseName string) (error, string, string) {
	if s.dbExists(databaseName) {
		return nil, uuid.NewV4().String(), uuid.NewV4().String()
	}
	return errors.New("DB does not exist"), "", ""
}

func (s *service) dbExists(databaseName string) bool {
	for _, db := range s.databases {
		if db == databaseName {
			return true
		}
	}
	return false
}
