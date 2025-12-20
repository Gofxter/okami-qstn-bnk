package models

import (
	"github.com/google/uuid"
	"okami-qstn-bnk/internal/pkg/types"
)

type InstantiateRequest struct {
	TemplateId uuid.UUID `json:"template_id"`
}

type (
	InstantiateResponse struct {
		SessionId uuid.UUID                       `json:"session_id"`
		Questions []QuestionInInstantiateResponse `json:"questions"`
	}

	QuestionInInstantiateResponse struct {
		QuestionId uuid.UUID                               `json:"question_id"`
		Type       types.QuestionType                      `json:"type" `
		Difficulty int                                     `json:"difficulty"`
		Text       string                                  `json:"text"`
		Options    []OptionInQuestionInInstantiateResponse `json:"options"`
	}

	OptionInQuestionInInstantiateResponse struct {
		Text      string `json:"text"`
		IsCorrect bool   `json:"is_correct"`
	}
)
