package repository

import (
	"go-micro-services-template/internal/app/port"
	postgressrepository "go-micro-services-template/internal/app/repository/postgres-repository"
)

func New(postgres *postgressrepository.PostgresRepository, misc port.MiscRepository, storage port.ObjectStorageRepository) *port.Repository {
	return &port.Repository{
		Sql:           postgres,
		Misc:          misc,
		ObjectStorage: storage,
	}
}
