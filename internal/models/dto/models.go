package models

import (
	"github.com/google/uuid"
	"okami-qstn-bnk/internal/pkg/types"
)

type Question struct {
	Id         uuid.UUID
	Role       types.Category
	Topic      string
	Type       types.QuestionType
	Difficulty int
	Text       string
}

type Option struct {
	Id         uuid.UUID
	QuestionId uuid.UUID
	Text       string
	IsCorrect  bool
}

type TestTemplate struct {
	Id      uuid.UUID
	Role    types.Category
	Purpose types.TemplatePurpose
}
