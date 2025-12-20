package storage

import (
	"context"
	"github.com/google/uuid"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type Storage interface {
	CreateQuestion(ctx context.Context, question *models.Question, options *[]models.Option) error
	GetQuestionById(ctx context.Context, id uuid.UUID) (*models.Question, error)
	GetQuestionsCollectionWithFilters(ctx context.Context, role *types.UserRole, topic *string, difficulty *int) ([]models.Question, error)
	UpdateQuestion(ctx context.Context, q models.Question) (*models.Question, error)
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	CreateTemplate(ctx context.Context, t models.TestTemplate) error
	GetTemplateById(ctx context.Context, id uuid.UUID) (*models.TestTemplate, error)
	GetTemplatesCollectionWithFilters(ctx context.Context, role *types.UserRole, purpose *types.TemplatePurpose) ([]models.TestTemplate, error)
	UpdateTemplate(ctx context.Context, t models.TestTemplate) (*models.TestTemplate, error)
	DeleteTemplate(ctx context.Context, id uuid.UUID) error
	GetRandomQuestion(ctx context.Context, templateId uuid.UUID) ([]models.Question, []models.Option, error)
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
}
