package models

import (
	"github.com/google/uuid"
	"okami-qstn-bnk/internal/pkg/types"
)

type (
	Question struct {
		Id         uuid.UUID       `json:"id"`
		Role       types.ModelRole `json:"role"`
		Topic      string          `json:"topic"`
		Type       types.ModelType `json:"type"`
		Options    *[]Option       `json:"options"`
		Difficulty int             `json:"difficulty"`
		Text       string          `json:"text"`
	}

	Option struct {
		Id         uuid.UUID `json:"id"`
		QuestionId uuid.UUID `json:"question_id"`
		Text       string    `json:"text"`
		IsCorrect  bool      `json:"is_correct"`
	}
)

type TestTemplate struct {
	Id      uuid.UUID          `json:"id"`
	Name    string             `json:"name"`
	Role    types.ModelRole    `json:"role"`
	Purpose types.ModelPurpose `json:"purpose"`
}

type TestTemplateQuestion struct {
	TemplateId uuid.UUID `json:"template_id"`
	QuestionId uuid.UUID `json:"question_id"`
	Order      int       `json:"order"`
}
