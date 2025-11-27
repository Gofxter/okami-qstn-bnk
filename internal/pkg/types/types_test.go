package types_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"okami-qstn-bnk/internal/pkg/types"
	"testing"
)

func TestRole(t *testing.T) {
	cases := []struct {
		name   string
		role   types.ModelRole
		result error
	}{
		{
			name:   "case without err",
			role:   "fronted_junior",
			result: nil,
		},
		{
			name:   "case without err",
			role:   "backend_junior",
			result: nil,
		},
		{
			name:   "case with err",
			role:   "qa_junior",
			result: errors.New(fmt.Sprintf("invalid type qa_junior")),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, types.ValidateRole(tc.role), tc.result)
		})
	}
}

func TestType(t *testing.T) {
	cases := []struct {
		name   string
		mType  types.ModelType
		result error
	}{
		{
			name:   "single choice happy case",
			mType:  "single_choice",
			result: nil,
		},
		{
			name:   "multi choice happy case",
			mType:  "multi_choice",
			result: nil,
		},
		{
			name:   "text happy case",
			mType:  "text",
			result: nil,
		},
		{
			name:   "case with err",
			mType:  "my_choice",
			result: errors.New("invalid type my_choice"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, types.ValidateType(tc.mType), tc.result)
		})
	}
}

func TestPurpose(t *testing.T) {
	cases := []struct {
		name    string
		purpose types.ModelPurpose
		result  error
	}{
		{
			name:    "case without err",
			purpose: "skills_assessment",
			result:  nil,
		},
		{
			name:    "case without err",
			purpose: "mock_interview",
			result:  nil,
		},
		{
			name:    "case with err",
			purpose: "another",
			result:  errors.New(fmt.Sprintf("invalid type another")),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, types.ValidatePurpose(tc.purpose), tc.result)
		})
	}
}
