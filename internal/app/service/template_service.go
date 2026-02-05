package service

import (
	"context"
)

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (s *TemplateService) Health(ctx context.Context) error {
	// logic จริง ๆ จะอยู่ตรงนี้
	return nil
}
