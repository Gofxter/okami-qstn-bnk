package service

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	"okami-qstn-bnk/internal/storage"
)

type QstnBnk struct {
	Logger  *zap.Logger
	Storage storage.Storage
}

func (q *QstnBnk) CreateQuestion(ctx context.Context, question *models.Question, options *[]models.Option) error {
	question.Id = uuid.New()
	if err := q.Storage.CreateQuestion(ctx, question, options); err != nil {
		return err
	}
	return nil
}

func (q *QstnBnk) GetQuestion(ctx context.Context, id uuid.UUID) (*models.Question, error) {
	result, err := q.Storage.GetQuestionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) GetQuestionsCollectionWithFilters(ctx context.Context, role *types.ModelRole, topic *string, difficulty *int) ([]models.Question, error) {
	result, err := q.Storage.GetQuestionsCollectionWithFilters(ctx, role, topic, difficulty)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) UpdateQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {
	result, err := q.Storage.UpdateQuestion(ctx, *question)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	if err := q.Storage.DeleteQuestion(ctx, id); err != nil {
		return err
	}

	return nil
}

func (q *QstnBnk) CreateTemplate(ctx context.Context, template *models.TestTemplate) error {
	template.Id = uuid.New()
	if err := q.Storage.CreateTemplate(ctx, *template); err != nil {
		return err
	}
	return nil
}

func (q *QstnBnk) GetTemplate(ctx context.Context, id uuid.UUID) (*models.TestTemplate, error) {
	result, err := q.Storage.GetTemplateById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) GetTemplatesCollectionWithFilters(ctx context.Context, role *types.ModelRole, purpose *types.ModelPurpose) ([]models.TestTemplate, error) {
	result, err := q.Storage.GetTemplatesCollectionWithFilters(ctx, role, purpose)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) UpdateTemplate(ctx context.Context, template *models.TestTemplate) (*models.TestTemplate, error) {
	result, err := q.Storage.UpdateTemplate(ctx, *template)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QstnBnk) DeleteTemplate(ctx context.Context, id uuid.UUID) error {
	if err := q.Storage.DeleteTemplate(ctx, id); err != nil {
		return err
	}

	return nil
}

func (q *QstnBnk) Instantiate(ctx context.Context, templateId uuid.UUID) (uuid.UUID, []models.Question, []models.Option, error) {
	var questions []models.Question
	var options []models.Option

	return uuid.New(), questions, options, nil
}
