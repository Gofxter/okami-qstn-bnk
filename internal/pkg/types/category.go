package types

import "strings"

type Category string

func (c Category) IsValid() bool {
	category := strings.Split(string(c), "_")

	var a bool
	switch category[0] {
	case "junior", "middle", "senior":
		a = true
	default:
		a = false
	}

	var b bool
	switch category[1] {
	case "qa", "ml", "backend", "frontend", "devops":
		b = true
	default:
		b = false
	}

	return a && b
}
