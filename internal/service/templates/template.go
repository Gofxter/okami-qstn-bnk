package templates

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	"okami-qstn-bnk/internal/storage"
)

type Template struct {
	Logger  *zap.Logger
	Storage storage.Storage
}

func (t *Template) CreateTemplate(ctx context.Context, template *models.TestTemplate) error {
	template.Id = uuid.New()
	t.Storage.CreateTemplate(ctx, *template)
	return nil
}

func (t *Template) GetTemplate(ctx context.Context, id uuid.UUID) (*models.TestTemplate, error) {
	result := t.Storage.GetTemplateById(ctx, id)
	return &result, nil
}

func (t *Template) GetTemplatesCollectionWithFilters(ctx context.Context, role *types.ModelRole, purpose *types.ModelPurpose) ([]models.TestTemplate, error) {
	result := t.Storage.GetTemplatesCollectionWithFilters(ctx, role, purpose)
	return result, nil
}

func (t *Template) UpdateTemplate(ctx context.Context, template *models.TestTemplate) (*models.TestTemplate, error) {
	result := t.Storage.UpdateTemplate(ctx, *template)
	return &result, nil
}

func (t *Template) DeleteTemplate(ctx context.Context, id uuid.UUID) error {
	t.Storage.DeleteTemplate(ctx, id)
	return nil
}
