package gorm

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type Gorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger, path string) *Gorm {
	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})

	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
		return nil
	}

	return &Gorm{
		db:     db,
		logger: logger,
	}
}

func (g *Gorm) CreateQuestion(q models.Question) {
	if err := g.db.Create(&q); err != nil {
		g.logger.Fatal("Failed to create question")
	}
}

func (g *Gorm) GetQuestionByID(id uuid.UUID) models.Question {
	var q models.Question
	g.db.Take(&q).Where("id = ?", id)
	return q
}

func (g *Gorm) GetQuestionsCollectionWithFilters(role *types.ModelRole, topic *string, difficulty *int) []models.Question {
	var qs []models.Question
	g.db.Find(&qs).Where("role = ?, topic = ?, difficulty = ?", role, topic, difficulty)

	return qs
}

func (g *Gorm) UpdateQuestion(q models.Question) models.Question {
	var result models.Question
	g.db.Model(result).Updates(q)

	g.db.Take(&result)
	return result
}

func (g *Gorm) DeleteQuestion(id uuid.UUID) {
	g.db.Delete(&models.Question{}, id)
}

func (g *Gorm) CreateTemplate(t models.TestTemplate) {
	if err := g.db.Create(&t); err != nil {
		g.logger.Fatal("Failed to create question")
	}
}

func (g *Gorm) GetTemplateById(id uuid.UUID) models.TestTemplate {
	var t models.TestTemplate
	g.db.Take(&t).Where("id = ?", id)
	return t
}

func (g *Gorm) GetTemplatesCollectionWithFilters(role *types.ModelRole, purpose *types.ModelPurpose) []models.TestTemplate {
	var qs []models.TestTemplate
	g.db.Find(&qs).Where("role = ?, purpose = ?", role, purpose)

	return qs
}

func (g *Gorm) UpdateTemplate(t models.TestTemplate) models.TestTemplate {
	var result models.TestTemplate
	g.db.Model(result).Updates(t)

	g.db.Take(&result)
	return result

}
func (g *Gorm) DeleteTemplate(id uuid.UUID) {
	g.db.Delete(&models.TestTemplate{}, id)
}
