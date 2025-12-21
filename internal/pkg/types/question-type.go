package types

type QuestionType string

const (
	SingleChoice QuestionType = "single_choice"
	MultiChoice  QuestionType = "multi_choice"
	Text         QuestionType = "text"
)

func (qt QuestionType) IsValid() bool {
	switch qt {
	case SingleChoice, MultiChoice, Text:
		return true
	default:
		return false
	}
}
