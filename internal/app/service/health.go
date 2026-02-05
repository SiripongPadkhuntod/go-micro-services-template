package service

import "context"

// Health validates downstream dependencies that the service relies on.
func (s *service) Health(ctx context.Context) error {
	// TODO: add dependency checks (database, message broker, etc.)
	return nil
}
