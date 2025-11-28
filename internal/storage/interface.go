package storage

import (
	"context"
	"github.com/google/uuid"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type Storage interface {
	CreateQuestion(ctx context.Context, q models.Question)
	GetQuestionByID(ctx context.Context, id uuid.UUID) models.Question
	GetQuestionsCollectionWithFilters(ctx context.Context, role *types.ModelRole, topic *string, difficulty *int) []models.Question
	UpdateQuestion(ctx context.Context, q models.Question) models.Question
	DeleteQuestion(ctx context.Context, id uuid.UUID)
	CreateTemplate(ctx context.Context, t models.TestTemplate)
	GetTemplateById(ctx context.Context, id uuid.UUID) models.TestTemplate
	GetTemplatesCollectionWithFilters(ctx context.Context, role *types.ModelRole, purpose *types.ModelPurpose) []models.TestTemplate
	UpdateTemplate(ctx context.Context, t models.TestTemplate) models.TestTemplate
	DeleteTemplate(ctx context.Context, id uuid.UUID)
}
