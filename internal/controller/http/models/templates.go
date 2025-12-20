package models

import (
	"github.com/google/uuid"
	models "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

type CreateTemplateRequest struct {
	Role    types.UserRole        `json:"role"`
	Purpose types.TemplatePurpose `json:"purpose"`
}

type GetTemplateByIDResponse struct {
	TemplateId uuid.UUID `json:"template_id"`
	QuestionId uuid.UUID `json:"question_id"`
	Order      int       `json:"order"`
}

type GetTemplatesWithFiltersResponse struct {
	Result []models.TestTemplate `json:"result"`
}

type (
	UpdateTemplateRequest struct {
		Role    types.UserRole        `json:"role"`
		Purpose types.TemplatePurpose `json:"purpose"`
	}
)
