package models

import (
	"okami-qstn-bnk/internal/pkg/types"
)

type QuestionsQueryRequest struct {
	Role       *types.ModelRole `query:"role"`
	Topic      *string          `query:"topic"`
	Difficulty *int             `query:"difficulty"`
}

type TemplatesQueryRequest struct {
	Role    *types.ModelRole    `query:"role"`
	Purpose *types.ModelPurpose `query:"purpose"`
}
