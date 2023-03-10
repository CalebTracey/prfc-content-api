package main

import (
	psql2 "github.com/calebtracey/api-template/internal/dao/psql"
	"github.com/calebtracey/api-template/internal/facade"
	"github.com/calebtracey/api-template/internal/facade/psql"
	config "github.com/calebtracey/config-yaml"
	log "github.com/sirupsen/logrus"
)

func initializeDAO(config config.Config) (facade.ServiceI, error) {
	psqlDbConfig, err := config.Database("PSQL")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return facade.Facade{
		PSQLDao: psql.Facade{
			PSQL: psql2.DAO{
				DB: psqlDbConfig.DB,
			},
			PSQLMapper: psql2.Mapper{},
		},
	}, nil
}
