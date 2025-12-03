package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	"okami-qstn-bnk/internal/service"
	mocks "okami-qstn-bnk/mocks/storage"
)

func newTestService(t *testing.T) (*service.QstnBnk, *mocks.MockStorage, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockStor := mocks.NewMockStorage(ctrl)
	logger, _ := zap.NewDevelopment()
	return &service.QstnBnk{Logger: logger, Storage: mockStor}, mockStor, ctrl
}

func TestQstnBnk_CreateQuestion(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	opts := &[]models.Option{{Text: "A", IsCorrect: true}}
	q := &models.Question{Role: "backend_junior", Topic: "go", Type: "text", Difficulty: 1, Text: "abc"}
	mock.EXPECT().CreateQuestion(ctx, gomock.Any(), opts).Return(nil)
	require.NoError(t, bnk.CreateQuestion(ctx, q, opts))

	mock.EXPECT().CreateQuestion(ctx, gomock.Any(), opts).Return(errors.New("fail"))
	require.Error(t, bnk.CreateQuestion(ctx, q, opts))
}

func TestQstnBnk_GetQuestion(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	id := uuid.New()
	q := &models.Question{Id: id}
	mock.EXPECT().GetQuestionById(ctx, id).Return(q, nil)
	got, err := bnk.GetQuestion(ctx, id)
	require.NoError(t, err)
	require.Equal(t, id, got.Id)

	mock.EXPECT().GetQuestionById(ctx, id).Return(nil, errors.New("fail"))
	_, err = bnk.GetQuestion(ctx, id)
	require.Error(t, err)
}

func TestQstnBnk_GetQuestionsCollectionWithFilters(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	r := types.ModelRole("backend_junior")
	topic := "go"
	diff := 1
	mock.EXPECT().GetQuestionsCollectionWithFilters(ctx, &r, &topic, &diff).Return([]models.Question{}, nil)
	l, err := bnk.GetQuestionsCollectionWithFilters(ctx, &r, &topic, &diff)
	require.NoError(t, err)
	require.NotNil(t, l)

	mock.EXPECT().GetQuestionsCollectionWithFilters(ctx, &r, &topic, &diff).Return(nil, errors.New("fail"))
	_, err = bnk.GetQuestionsCollectionWithFilters(ctx, &r, &topic, &diff)
	require.Error(t, err)
}

func TestQstnBnk_UpdateQuestion(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	q := &models.Question{Id: uuid.New()}
	mock.EXPECT().UpdateQuestion(ctx, *q).Return(q, nil)
	upd, err := bnk.UpdateQuestion(ctx, q)
	require.NoError(t, err)
	require.Equal(t, q.Id, upd.Id)

	mock.EXPECT().UpdateQuestion(ctx, *q).Return(nil, errors.New("fail"))
	_, err = bnk.UpdateQuestion(ctx, q)
	require.Error(t, err)
}

func TestQstnBnk_DeleteQuestion(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	id := uuid.New()
	mock.EXPECT().DeleteQuestion(ctx, id).Return(nil)
	require.NoError(t, bnk.DeleteQuestion(ctx, id))
	mock.EXPECT().DeleteQuestion(ctx, id).Return(errors.New("fail"))
	require.Error(t, bnk.DeleteQuestion(ctx, id))
}

func TestQstnBnk_CreateTemplate(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	temp := &models.TestTemplate{Role: "backend_junior", Purpose: "mock_interview"}
	mock.EXPECT().CreateTemplate(ctx, gomock.Any()).Return(nil)
	require.NoError(t, bnk.CreateTemplate(ctx, temp))
	mock.EXPECT().CreateTemplate(ctx, gomock.Any()).Return(errors.New("fail"))
	require.Error(t, bnk.CreateTemplate(ctx, temp))
}

func TestQstnBnk_GetTemplate(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	id := uuid.New()
	tmpl := &models.TestTemplate{Id: id}
	mock.EXPECT().GetTemplateById(ctx, id).Return(tmpl, nil)
	got, err := bnk.GetTemplate(ctx, id)
	require.NoError(t, err)
	require.Equal(t, id, got.Id)

	mock.EXPECT().GetTemplateById(ctx, id).Return(nil, errors.New("fail"))
	_, err = bnk.GetTemplate(ctx, id)
	require.Error(t, err)
}

func TestQstnBnk_GetTemplatesCollectionWithFilters(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	role := types.ModelRole("backend_junior")
	purp := types.ModelPurpose("mock_interview")
	mock.EXPECT().GetTemplatesCollectionWithFilters(ctx, &role, &purp).Return([]models.TestTemplate{}, nil)
	l, err := bnk.GetTemplatesCollectionWithFilters(ctx, &role, &purp)
	require.NoError(t, err)
	require.NotNil(t, l)

	mock.EXPECT().GetTemplatesCollectionWithFilters(ctx, &role, &purp).Return(nil, errors.New("fail"))
	_, err = bnk.GetTemplatesCollectionWithFilters(ctx, &role, &purp)
	require.Error(t, err)
}

func TestQstnBnk_UpdateTemplate(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	tmpl := &models.TestTemplate{Id: uuid.New()}
	mock.EXPECT().UpdateTemplate(ctx, *tmpl).Return(tmpl, nil)
	upd, err := bnk.UpdateTemplate(ctx, tmpl)
	require.NoError(t, err)
	require.Equal(t, tmpl.Id, upd.Id)

	mock.EXPECT().UpdateTemplate(ctx, *tmpl).Return(nil, errors.New("fail"))
	_, err = bnk.UpdateTemplate(ctx, tmpl)
	require.Error(t, err)
}

func TestQstnBnk_DeleteTemplate(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	id := uuid.New()
	mock.EXPECT().DeleteTemplate(ctx, id).Return(nil)
	require.NoError(t, bnk.DeleteTemplate(ctx, id))
	mock.EXPECT().DeleteTemplate(ctx, id).Return(errors.New("fail"))
	require.Error(t, bnk.DeleteTemplate(ctx, id))
}

func TestQstnBnk_Instantiate(t *testing.T) {
	bnk, mock, ctrl := newTestService(t)
	defer ctrl.Finish()
	ctx := context.Background()
	tplId := uuid.New()
	qq := []models.Question{{Id: uuid.New()}}
	opts := []models.Option{{}}
	mock.EXPECT().GetRandomQuestion(ctx, tplId).Return(qq, opts, nil)
	sid, gotqq, gotopts, err := bnk.Instantiate(ctx, tplId)
	require.NoError(t, err)
	require.Equal(t, qq, gotqq)
	require.Equal(t, opts, gotopts)
	require.NotEqual(t, uuid.Nil, sid)

	mock.EXPECT().GetRandomQuestion(ctx, tplId).Return(nil, nil, errors.New("fail"))
	_, _, _, err = bnk.Instantiate(ctx, tplId)
	require.Error(t, err)
}
