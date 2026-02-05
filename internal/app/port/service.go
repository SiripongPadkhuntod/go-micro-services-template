package port

import "context"

type Service interface {
	Health(ctx context.Context) error
	
}
