package types

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
