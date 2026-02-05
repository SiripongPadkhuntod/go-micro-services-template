package handler

import "go-micro-services-template/internal/app/port"

type Handler struct {
	svc port.Service
}

func New(svc port.Service) port.Handler {
	return &Handler{svc: svc}
}
