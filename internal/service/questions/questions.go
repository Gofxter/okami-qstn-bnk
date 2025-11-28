package questions

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	"okami-qstn-bnk/internal/storage"
)

type Question struct {
	Logger  *zap.Logger
	Storage storage.Storage
}

func (q *Question) CreateQuestion(ctx context.Context, question *models.Question) error {
	question.Id = uuid.New()
	q.Storage.CreateQuestion(ctx, *question)
	return nil
}

func (q *Question) GetQuestion(ctx context.Context, id uuid.UUID) (*models.Question, error) {
	result := q.Storage.GetQuestionById(ctx, id)
	return &result, nil
}

func (q *Question) GetQuestionsCollectionWithFilters(ctx context.Context, role *types.ModelRole, topic *string, difficulty *int) ([]models.Question, error) {
	result := q.Storage.GetQuestionsCollectionWithFilters(ctx, role, topic, difficulty)
	return result, nil
}

func (q *Question) UpdateQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {
	result := q.Storage.UpdateQuestion(ctx, *question)
	return &result, nil
}

func (q *Question) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	q.Storage.DeleteQuestion(ctx, id)
	return nil
}
