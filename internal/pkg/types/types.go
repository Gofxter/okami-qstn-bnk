package types

import (
	"errors"
	"fmt"
)

type (
	ModelRole    string
	ModelType    string
	ModelPurpose string
)

func ValidateRole(r ModelRole) error {
	switch r {
	case "frontend_junior":
		return nil
	case "backend_junior":
		return nil
	}

	return errors.New(fmt.Sprintf("invalid type %s", r))
}

func ValidateType(t ModelType) error {
	switch t {
	case "single_choice":
		return nil
	case "multi_choice":
		return nil
	case "text":
		return nil
	}

	return errors.New(fmt.Sprintf("invalid type %s", t))
}

func ValidatePurpose(p ModelPurpose) error {
	switch p {
	case "skills_assessment":
		return nil
	case "mock_interview":
		return nil
	}

	return errors.New(fmt.Sprintf("invalid type %s", p))
}
