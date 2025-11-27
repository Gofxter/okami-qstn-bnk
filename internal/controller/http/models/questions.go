package models

import (
	"github.com/google/uuid"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type (
	CreateQuestionRequest struct {
		Role       types.ModelRole                `json:"role"`
		Topic      string                         `json:"topic"`
		Type       types.ModelType                `json:"type"`
		Options    *[]OptionCreateQuestionRequest `json:"options"`
		Difficulty int                            `json:"difficulty"`
		Text       string                         `json:"text"`
	}

	OptionCreateQuestionRequest struct {
		Text      string `json:"text"`
		IsCorrect bool   `json:"is_correct"`
	}
)

type (
	GetQuestionByIDResponse struct {
		Id         uuid.UUID                                 `json:"id"`
		Role       types.ModelRole                           `json:"role"`
		Topic      string                                    `json:"topic"`
		Type       types.ModelType                           `json:"type"`
		Options    *[]OptionGetQuestionsByQuestionIDResponse `json:"options"`
		Difficulty int                                       `json:"difficulty"`
		Text       string                                    `json:"text"`
	}

	OptionGetQuestionsByQuestionIDResponse struct {
		Id         uuid.UUID `json:"id"`
		QuestionId uuid.UUID `json:"question_id"`
		Text       string    `json:"text"`
		IsCorrect  bool      `json:"is_correct"`
	}
)

type GetQuestionsWithFiltersResponse struct {
	Result []models.Question `json:"result"`
}

type (
	UpdateQuestionRequest struct {
		Role       *types.ModelRole               `json:"role"`
		Topic      *string                        `json:"topic"`
		Type       *types.ModelType               `json:"type"`
		Options    *[]OptionUpdateQuestionRequest `json:"options"`
		Difficulty *int                           `json:"difficulty"`
		Text       *string                        `json:"text"`
	}

	OptionUpdateQuestionRequest struct {
		Text *string `json:"text"`
	}
	UpdateQuestionResponse struct {
		Id         uuid.UUID                       `json:"id"`
		Role       types.ModelRole                 `json:"role"`
		Topic      string                          `json:"topic"`
		Type       types.ModelType                 `json:"type"`
		Options    *[]OptionUpdateQuestionResponse `json:"options"`
		Difficulty int                             `json:"difficulty"`
		Text       string                          `json:"text"`
	}

	OptionUpdateQuestionResponse struct {
		Id         uuid.UUID `json:"id"`
		QuestionId uuid.UUID `json:"question_id"`
		Text       string    `json:"text"`
		IsCorrect  bool      `json:"is_correct"`
	}
)
