package main

// default

// imports

// APIPackage          string = restapi
// Package             string = operations
// ReceiverName        string = o
// Name                string = SearchkitApis
// Principal           string = interface{}
// DefaultConsumes     string = application/json
// DefaultProduces     string = application/json
// Host                string = localhost:5000
// BasePath            string =    /

// "github.com/roscopecoltran/admin-on-rest-server/server/database"
// "github.com/roscopecoltran/admin-on-rest-server/server/models"

type SearchkitApisService struct {
	Health bool
}

func (s *SearchkitApisService) Healthy() bool {
	return s.Health
}
