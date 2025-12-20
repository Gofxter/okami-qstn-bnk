package models

import (
	"okami-qstn-bnk/internal/pkg/types"
)

type QuestionsQueryRequest struct {
	Role       *types.UserRole `query:"role"`
	Topic      *string         `query:"topic"`
	Difficulty *int            `query:"difficulty"`
}

type TemplatesQueryRequest struct {
	Role    *types.UserRole        `query:"role"`
	Purpose *types.TemplatePurpose `query:"purpose"`
}
