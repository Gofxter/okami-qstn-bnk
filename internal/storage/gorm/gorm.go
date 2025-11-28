package gorm

import (
	"context"
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

func (g *Gorm) CreateQuestion(ctx context.Context, q models.Question) {
	g.db.WithContext(ctx).Create(&q)
}

func (g *Gorm) GetQuestionById(ctx context.Context, id uuid.UUID) models.Question {
	var q models.Question
	g.db.WithContext(ctx).Take(&q).Where("id = ?", id)
	return q
}

func (g *Gorm) GetQuestionsCollectionWithFilters(ctx context.Context, role *types.ModelRole, topic *string, difficulty *int) []models.Question {
	var qs []models.Question
	g.db.WithContext(ctx).Find(&qs).Where("role = ?, topic = ?, difficulty = ?", role, topic, difficulty)

	return qs
}

func (g *Gorm) UpdateQuestion(ctx context.Context, q models.Question) models.Question {
	var result models.Question
	g.db.WithContext(ctx).Model(result).Updates(q)

	g.db.WithContext(ctx).Take(&result)
	return result
}

func (g *Gorm) DeleteQuestion(ctx context.Context, id uuid.UUID) {
	g.db.WithContext(ctx).Delete(&models.Question{}, id)
}

func (g *Gorm) CreateTemplate(ctx context.Context, t models.TestTemplate) {
	g.db.WithContext(ctx).Create(&t)
}

func (g *Gorm) GetTemplateById(ctx context.Context, id uuid.UUID) models.TestTemplate {
	var t models.TestTemplate
	g.db.WithContext(ctx).First(&t).Where("id = ?", id)
	return t
}

func (g *Gorm) GetTemplatesCollectionWithFilters(ctx context.Context, role *types.ModelRole, purpose *types.ModelPurpose) []models.TestTemplate {
	var qs []models.TestTemplate
	g.db.WithContext(ctx).Find(&qs).Where("role = ?, purpose = ?", role, purpose)

	return qs
}

func (g *Gorm) UpdateTemplate(ctx context.Context, t models.TestTemplate) models.TestTemplate {
	var result models.TestTemplate
	g.db.WithContext(ctx).Model(result).Where("id = ?", t.Id).Updates(t)

	g.db.WithContext(ctx).Take(&result)
	return result

}
func (g *Gorm) DeleteTemplate(ctx context.Context, id uuid.UUID) {
	g.db.WithContext(ctx).Delete(&models.TestTemplate{}, id)
}

func (g *Gorm) Ping(ctx context.Context) error {
	storage, err := g.db.WithContext(ctx).DB()
	if err != nil {
		g.logger.Error("Failed to connect to database", zap.Error(err))
		return err
	}

	if err := storage.Ping(); err != nil {
		g.logger.Error("Failed to ping to database", zap.Error(err))
		return err
	}

	return nil
}

func (g *Gorm) Close(ctx context.Context) error {
	storage, err := g.db.WithContext(ctx).DB()
	if err != nil {
		g.logger.Error("Failed to connect to database", zap.Error(err))
		return err
	}

	if err := storage.Close(); err != nil {
		g.logger.Error("Failed to close database", zap.Error(err))
		return err
	}

	return nil
}
