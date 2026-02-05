package server

import (
	"context"
	"go-micro-services-template/internal/app/handler"
	"go-micro-services-template/internal/property"
)

func RunServer(ctx context.Context) error {
	// --- Initialize logger ---

	// --- Initialize dependencies ---
	adpt := initInfrastructure(ctx)
	repo := initRepository(ctx)
	svc := initService(*adpt, *repo)
	h := handler.New(svc)

	// --- Initialize HTTP server ---
	engine := initServer(h)

	// --- Run server ---
	return RunServerWithContext(ctx, engine)
}

func RunServerWithConfig(cfg *property.Config) error {
	ctx := context.Background()
	// --- Initialize logger ---

	// --- Initialize dependencies ---
	adpt := initInfrastructure(ctx)
	repo := initRepository(ctx)
	svc := initService(*adpt, *repo)
	h := handler.New(svc)

	// --- Initialize HTTP server ---
	engine := initServer(h)

	// --- Run server ---
	return RunServerWithContext(ctx, engine)
}
