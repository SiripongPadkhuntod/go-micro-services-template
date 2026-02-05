package service

import (
	"go-micro-services-template/internal/app/port"
)

type service struct {
	adapt port.Adapter
	repo port.Repository
}

func New(adpt port.Adapter, repo port.Repository) port.Service {
	return &service{
		adapt: adpt,
		repo: repo,
	}
}
