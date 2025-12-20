package gorm

import (
	"context"
	"errors"
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

func (g *Gorm) CreateQuestion(ctx context.Context, q *models.Question, options *[]models.Option) error {
	if err := g.db.WithContext(ctx).Create(&q).Error; err != nil {
		g.logger.Error("Failed to create question", zap.Error(err))
		return err
	}

	if options != nil {
		for index := range *options {
			(*options)[index].Id = uuid.New()
			(*options)[index].QuestionId = q.Id

			if err := g.db.WithContext(ctx).Create(&(*options)[index]).Error; err != nil {
				g.logger.Error("Failed to create option", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func (g *Gorm) GetQuestionById(ctx context.Context, id uuid.UUID) (*models.Question, error) {
	var q models.Question

	if err := g.db.WithContext(ctx).Where("id = ?", id).First(&q).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			g.logger.Warn("Question not found", zap.String("id", id.String()))
			return nil, errors.New("question not found")
		}

		g.logger.Error("Failed to get question by id", zap.Error(err))
		return nil, err
	}

	return &q, nil
}

func (g *Gorm) GetQuestionsCollectionWithFilters(ctx context.Context, role *types.UserRole, topic *string, difficulty *int) ([]models.Question, error) {
	var qs []models.Question

	query := g.db.WithContext(ctx).Model(&models.Question{})

	if role != nil {
		query = query.Where("role = ?", *role)
	}
	if topic != nil {
		query = query.Where("topic = ?", *topic)
	}
	if difficulty != nil {
		query = query.Where("difficulty = ?", *difficulty)
	}

	err := query.Find(&qs).Error
	if err != nil {
		g.logger.Error("Failed to get questions with filters", zap.Error(err))
		return nil, err
	}

	return qs, nil
}

func (g *Gorm) UpdateQuestion(ctx context.Context, q models.Question) (*models.Question, error) {
	result := g.db.WithContext(ctx).Model(&models.Question{}).Where("id = ?", q.Id).Updates(q)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			g.logger.Warn("Question not found", zap.Any("id", q.Id))
			return nil, errors.New("question not found")
		}

		g.logger.Error("Failed to update question", zap.Error(result.Error))
		return nil, result.Error
	}

	var updatedTemplate *models.Question
	err := g.db.WithContext(ctx).Where("id = ?", q.Id).First(&updatedTemplate).Error
	if err != nil {
		g.logger.Error("Failed to fetch question", zap.Error(err))
		return nil, err
	}

	return updatedTemplate, nil
}

func (g *Gorm) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	option := g.db.WithContext(ctx).Where("question_id = ?", id).Delete(&models.Option{})
	if option.Error != nil {
		if errors.Is(option.Error, gorm.ErrRecordNotFound) {
			g.logger.Warn("Question not found", zap.String("id", id.String()))
			return errors.New("questions option not found")
		}

		g.logger.Error("Failed to delete questions option", zap.Error(option.Error))
		return option.Error
	}

	question := g.db.WithContext(ctx).Delete(&models.Question{}, id)
	if question.Error != nil {
		if errors.Is(question.Error, gorm.ErrRecordNotFound) {
			g.logger.Warn("Question not found", zap.String("id", id.String()))
			return errors.New("question not found")
		}

		g.logger.Error("Failed to delete question", zap.Error(question.Error))
		return question.Error
	}

	return nil
}

func (g *Gorm) CreateTemplate(ctx context.Context, t models.TestTemplate) error {
	if err := g.db.WithContext(ctx).Create(&t).Error; err != nil {
		g.logger.Error("Failed to create question", zap.Error(err))
		return err
	}
	return nil
}

func (g *Gorm) GetTemplateById(ctx context.Context, id uuid.UUID) (*models.TestTemplate, error) {
	var t models.TestTemplate

	if err := g.db.WithContext(ctx).Where("id = ?", id).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			g.logger.Warn("Template not found", zap.String("id", id.String()))
			return nil, errors.New("template not found")
		}

		g.logger.Error("Failed to get template by id", zap.Error(err))
		return nil, err
	}

	return &t, nil
}

func (g *Gorm) GetTemplatesCollectionWithFilters(ctx context.Context, role *types.UserRole, purpose *types.TemplatePurpose) ([]models.TestTemplate, error) {
	var qs []models.TestTemplate

	query := g.db.WithContext(ctx).Model(&models.TestTemplate{})

	if role != nil {
		query = query.Where("role = ?", *role)
	}

	if purpose != nil {
		query = query.Where("purpose = ?", *purpose)
	}

	if err := query.Find(&qs).Error; err != nil {
		g.logger.Error("Failed to get templates with filters", zap.Error(err))
		return nil, err
	}

	return qs, nil
}

func (g *Gorm) UpdateTemplate(ctx context.Context, t models.TestTemplate) (*models.TestTemplate, error) {
	result := g.db.WithContext(ctx).Model(&models.TestTemplate{}).Where("id = ?", t.Id).Updates(t)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			g.logger.Warn("Template not found", zap.Any("id", t.Id))
			return nil, errors.New("template not found")
		}

		g.logger.Error("Failed to update template", zap.Error(result.Error))
		return nil, result.Error
	}

	var updatedTemplate *models.TestTemplate
	err := g.db.WithContext(ctx).Where("id = ?", t.Id).First(&updatedTemplate).Error
	if err != nil {
		g.logger.Error("Failed to fetch updated template", zap.Error(err))
		return nil, err
	}

	return updatedTemplate, nil

}

func (g *Gorm) DeleteTemplate(ctx context.Context, id uuid.UUID) error {
	result := g.db.WithContext(ctx).Delete(&models.TestTemplate{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			g.logger.Warn("Template not found", zap.String("id", id.String()))
			return errors.New("template not found")
		}

		g.logger.Error("Failed to delete template", zap.Error(result.Error))
		return result.Error
	}

	return nil
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

func (g *Gorm) GetRandomQuestion(ctx context.Context, templateId uuid.UUID) ([]models.Question, []models.Option, error) {
	var qs []models.Question
	var opt []models.Option
	var template models.TestTemplate
	var questionIDs []uuid.UUID

	if err := g.db.WithContext(ctx).First(&template, "id = ?", templateId).Error; err != nil {
		g.logger.Error("Failed to get template", zap.Error(err))
		return nil, nil, err
	}

	if err := g.db.Order("RANDOM()").Limit(3).Where("role = ?", template.Role).Find(&qs).Error; err != nil {
		g.logger.Error("Failed to get questions", zap.Error(err))
		return nil, nil, err
	}

	for _, q := range qs {
		questionIDs = append(questionIDs, q.Id)
	}

	if err := g.db.Where("question_id IN ?", questionIDs).Find(&opt).Error; err != nil {
		g.logger.Error("Failed to get options", zap.Error(err))
		return qs, nil, err
	}

	return qs, opt, nil
}
