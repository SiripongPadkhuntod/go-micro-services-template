package postgresrepository

import "context"

type PostgresRepository struct {
}

func New(ctx context.Context) *PostgresRepository {
	return &PostgresRepository{}
}
