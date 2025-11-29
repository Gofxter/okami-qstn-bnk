package service

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	"okami-qstn-bnk/internal/storage"
)

type Service interface {
	CreateQuestion(ctx context.Context, question *models.Question, options *[]models.Option) error
	GetQuestion(ctx context.Context, id uuid.UUID) (*models.Question, error)
	GetQuestionsCollectionWithFilters(ctx context.Context, role *types.ModelRole, topic *string, difficulty *int) ([]models.Question, error)
	UpdateQuestion(ctx context.Context, question *models.Question) (*models.Question, error)
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	CreateTemplate(ctx context.Context, template *models.TestTemplate) error
	GetTemplate(ctx context.Context, id uuid.UUID) (*models.TestTemplate, error)
	GetTemplatesCollectionWithFilters(ctx context.Context, role *types.ModelRole, purpose *types.ModelPurpose) ([]models.TestTemplate, error)
	UpdateTemplate(ctx context.Context, template *models.TestTemplate) (*models.TestTemplate, error)
	DeleteTemplate(ctx context.Context, id uuid.UUID) error
}

func RegisterServices(logger *zap.Logger, storage storage.Storage) *QstnBnk {
	return &QstnBnk{
		Logger:  logger,
		Storage: storage,
	}
}
