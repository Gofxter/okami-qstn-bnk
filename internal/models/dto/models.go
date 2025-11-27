package models

import (
	"github.com/google/uuid"
	"okami-qstn-bnk/internal/pkg/types"
)

type (
	Question struct {
		Id         uuid.UUID       `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique"`
		Role       types.ModelRole `json:"role" gorm:"type:varchar(25);not null"`
		Topic      string          `json:"topic" gorm:"type:varchar(150);not null"`
		Type       types.ModelType `json:"type" gorm:"type:varchar(25);not null"`
		Options    *[]Option       `json:"options"`
		Difficulty int             `json:"difficulty" gorm:"type:smallint;not null"`
		Text       string          `json:"text" gorm:"type:text;not null"`
	}

	Option struct {
		Id         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique;not null"`
		QuestionId uuid.UUID `json:"question_id" gorm:"type:uuid;not null"`
		Text       string    `json:"text" gorm:"type:text;not null"`
		IsCorrect  bool      `json:"is_correct" gorm:"type:boolean;not null"`
	}
)

type TestTemplate struct {
	Id      uuid.UUID          `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique"`
	Role    types.ModelRole    `json:"role" gorm:"type:varchar(25);not null"`
	Purpose types.ModelPurpose `json:"purpose" gorm:"type:varchar(50);not null"`
}

type TestTemplateQuestion struct {
	TemplateId uuid.UUID `json:"template_id"`
	QuestionId uuid.UUID `json:"question_id"`
	Order      int       `json:"order"`
}
