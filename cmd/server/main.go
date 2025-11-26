package main

import (
	"fmt"
	"github.com/google/uuid"
)

type S struct {
	a uuid.UUID // default 0
	b string
}

func main() {
	var s = S{b: "2"}
	ChangeA(&s)
	fmt.Println(s)
}

func ChangeA(s *S) {
	s.a = uuid.New()
}
