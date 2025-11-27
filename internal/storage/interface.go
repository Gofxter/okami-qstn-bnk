package storage

import (
	"github.com/google/uuid"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type Storage interface {
	CreateQuestion(q models.Question)
	GetQuestionByID(id uuid.UUID) models.Question
	GetQuestionsCollectionWithFilters(role *types.ModelRole, topic *string, difficulty *int) []models.Question
	UpdateQuestion(q models.Question) models.Question
	DeleteQuestion(id uuid.UUID)
	CreateTemplate(t models.TestTemplate)
	GetTemplateById(id uuid.UUID) models.TestTemplate
	GetTemplatesCollectionWithFilters(role *types.ModelRole, purpose *types.ModelPurpose) []models.TestTemplate
	UpdateTemplate(t models.TestTemplate) models.TestTemplate
	DeleteTemplate(id uuid.UUID)
}
