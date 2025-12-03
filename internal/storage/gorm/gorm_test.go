package gorm_test

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	cfg "okami-qstn-bnk/internal/config"
	dto "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
	gormdb "okami-qstn-bnk/internal/storage/gorm"
)

func getConfigPath(t *testing.T) string {
	root := getRepoRoot(t)
	return filepath.Join(root, "config", "config.yaml")
}

func getRepoRoot(t *testing.T) string {
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "Не удалось определить путь тестового файла")
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", "..", ".."))
}

func getTestDSN(t *testing.T) string {
	logger := newLogger()
	config := cfg.LoadConfig(getConfigPath(t), logger)
	if &config == nil {
		require.FailNow(t, "Не удалось загрузить config.yaml для тестов базы данных")
	}

	config.Storage.SetURI(logger)
	dsn := config.Storage.GetURI()
	if dsn == "" {
		require.FailNow(t, "Строка подключения к БД не задана после парсинга config.yaml")
	}
	return dsn
}

func newTestDB(t *testing.T) (*gorm.DB, func()) {
	dsn := getTestDSN(t)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Поднимаем схему (создаём таблицы, если их ещё нет)
	err = db.AutoMigrate(&dto.Question{}, &dto.Option{}, &dto.TestTemplate{})
	require.NoError(t, err)

	// Чистим данные перед каждым тестом
	db.Exec("TRUNCATE TABLE options, questions, test_templates RESTART IDENTITY CASCADE")
	return db, func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func TestGorm_Question_CRUD(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	q := &dto.Question{
		Id:         uuid.New(),
		Role:       "backend_junior",
		Topic:      "go-sql",
		Type:       "single_choice",
		Difficulty: 3,
		Text:       "SQL DB?",
	}
	options := []dto.Option{{Text: "SQLite", IsCorrect: true}, {Text: "MySQL", IsCorrect: false}}

	t.Run("create & get", func(t *testing.T) {
		err := g.CreateQuestion(ctx, q, &options)
		require.NoError(t, err)

		q2, err := g.GetQuestionById(ctx, q.Id)
		require.NoError(t, err)
		require.Equal(t, q.Id, q2.Id)
		qs, err := g.GetQuestionsCollectionWithFilters(ctx, &q.Role, &q.Topic, &q.Difficulty)
		require.NoError(t, err)
		require.Len(t, qs, 1)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := g.GetQuestionById(ctx, uuid.New())
		require.Error(t, err)
	})

	t.Run("update", func(t *testing.T) {
		q.Text = "Updated?"
		upd, err := g.UpdateQuestion(ctx, *q)
		require.NoError(t, err)
		require.Equal(t, "Updated?", upd.Text)
	})

	t.Run("delete", func(t *testing.T) {
		err := g.DeleteQuestion(ctx, q.Id)
		require.NoError(t, err)
		_, err = g.GetQuestionById(ctx, q.Id)
		require.Error(t, err)
	})
}

func TestGorm_Question_Filters(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	q1 := &dto.Question{Id: uuid.New(), Role: "backend_junior", Topic: "sql", Type: "single_choice", Difficulty: 4, Text: "q1"}
	q2 := &dto.Question{Id: uuid.New(), Role: "frontend_junior", Topic: "html", Type: "text", Difficulty: 1, Text: "q2"}
	_ = g.CreateQuestion(ctx, q1, nil)
	_ = g.CreateQuestion(ctx, q2, nil)
	r := types.ModelRole("backend_junior")
	d, tpc := 4, "sql"

	qs, err := g.GetQuestionsCollectionWithFilters(ctx, &r, &tpc, &d)
	require.NoError(t, err)
	require.Len(t, qs, 1)
	qs, err = g.GetQuestionsCollectionWithFilters(ctx, nil, nil, nil)
	require.NoError(t, err)
	require.Len(t, qs, 2)
}

func TestGorm_Question_NotFound_Update_Delete(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	r := dto.Question{Id: uuid.New()}
	_, err := g.UpdateQuestion(ctx, r)
	require.Error(t, err)
	err = g.DeleteQuestion(ctx, r.Id)
	require.NoError(t, err) // delete ошибка не если не было опций и вопроса
}

func TestGorm_Template_CRUD(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	template := dto.TestTemplate{Id: uuid.New(), Role: "backend_junior", Purpose: "skills_assessment"}
	err := g.CreateTemplate(ctx, template)
	require.NoError(t, err)

	got, err := g.GetTemplateById(ctx, template.Id)
	require.NoError(t, err)
	require.Equal(t, template.Id, got.Id)

	// Update
	template.Purpose = "mock_interview"
	upd, err := g.UpdateTemplate(ctx, template)
	require.NoError(t, err)
	require.Equal(t, "mock_interview", string(upd.Purpose))

	// Filter
	r := types.ModelRole("backend_junior")
	p := types.ModelPurpose("mock_interview")
	lst, err := g.GetTemplatesCollectionWithFilters(ctx, &r, &p)
	require.NoError(t, err)
	require.Len(t, lst, 1)

	// Delete
	err = g.DeleteTemplate(ctx, template.Id)
	require.NoError(t, err)
	_, err = g.GetTemplateById(ctx, template.Id)
	require.Error(t, err)
}

func TestGorm_PingClose(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	require.NoError(t, g.Ping(ctx))
	require.NoError(t, g.Close(ctx))
}

func TestGorm_GetRandomQuestion(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	tpl := dto.TestTemplate{Id: uuid.New(), Role: "backend_junior", Purpose: "skills_assessment"}
	_ = g.CreateTemplate(ctx, tpl)
	for i := 0; i < 5; i++ {
		q := &dto.Question{
			Id:         uuid.New(),
			Role:       tpl.Role,
			Topic:      "any",
			Type:       "single_choice",
			Difficulty: 2,
			Text:       "RandomQ?",
		}
		_ = g.CreateQuestion(ctx, q, nil)
	}
	qs, opts, err := g.GetRandomQuestion(ctx, tpl.Id)
	require.NoError(t, err)
	require.LessOrEqual(t, len(qs), 3)
	_ = opts // могут быть пусты без опций
}

func TestGorm_Errors_GetRandomQuestion(t *testing.T) {
	ctx := context.Background()
	_, cleanup := newTestDB(t)
	defer cleanup()
	logger := newLogger()
	dsn := getTestDSN(t)
	g := gormdb.NewStorage(logger, dsn)

	_, _, err := g.GetRandomQuestion(ctx, uuid.New())
	require.Error(t, err)
}
