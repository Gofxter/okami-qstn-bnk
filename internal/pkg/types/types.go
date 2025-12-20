package types

type UserRole string

const (
	UserRoleNone  UserRole = ""
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

func (ur UserRole) IsValid() bool {
	switch ur {
	case UserRoleNone, UserRoleAdmin, UserRoleUser:
		return true
	}
	return false
}

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

type TemplatePurpose string

const (
	SkillsAssessment TemplatePurpose = "skills_assessment"
	MockInterview    TemplatePurpose = "mock_interview"
)

func (pt TemplatePurpose) IsValid() bool {
	switch pt {
	case SkillsAssessment, MockInterview:
		return true
	default:
		return false
	}
}
