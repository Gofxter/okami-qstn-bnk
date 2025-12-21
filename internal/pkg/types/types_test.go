package types_test

import (
	"github.com/stretchr/testify/assert"
	"okami-qstn-bnk/internal/pkg/types"
	"testing"
)

func TestCategory(t *testing.T) {
	cases := []struct {
		name   string
		mType  types.Category
		result bool
	}{
		{
			name:   "case without err",
			mType:  "junior_frontend",
			result: true,
		},
		{
			name:   "case without err",
			mType:  "junior_backend",
			result: true,
		},
		{
			name:   "case with err",
			mType:  "ultra_manager",
			result: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.mType.IsValid(), tc.result)
		})
	}
}

func TestType(t *testing.T) {
	cases := []struct {
		name   string
		mType  types.QuestionType
		result bool
	}{
		{
			name:   "single choice happy case",
			mType:  "single_choice",
			result: true,
		},
		{
			name:   "multi choice happy case",
			mType:  "multi_choice",
			result: true,
		},
		{
			name:   "text happy case",
			mType:  "text",
			result: true,
		},
		{
			name:   "case with err",
			mType:  "my_choice",
			result: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.mType.IsValid(), tc.result)
		})
	}
}

func TestPurpose(t *testing.T) {
	cases := []struct {
		name   string
		mType  types.TemplatePurpose
		result bool
	}{
		{
			name:   "case without err",
			mType:  "skills_assessment",
			result: true,
		},
		{
			name:   "case without err",
			mType:  "mock_interview",
			result: true,
		},
		{
			name:   "case with err",
			mType:  "another",
			result: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.mType.IsValid(), tc.result)
		})
	}
}
