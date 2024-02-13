package storage

import (
	"github.com/agustinrabini/api-items-project/cmd/api/config"

	"github.com/agustinrabini/go-toolkit/gonosql"
)

func NewNoSQL() *gonosql.Data {
	config := getDBConfig()
	return gonosql.NewNoSQL(config)
}

func getDBConfig() gonosql.Config {
	return gonosql.Config{
		Username: config.ConfMap.MongoUser,
		Password: config.ConfMap.MongoPassword,
		Host:     config.ConfMap.MongoHost,
		Database: config.ConfMap.MongoDataBase,
	}
}
