package strategy

import "strings"

type Strategy interface {
	Search(boxes []int, prisonerNumber int) bool
}

func Get(name string) Strategy {
	switch strings.ToLower(name) {
	case "loop":
		return &Loop{}
	case "random":
		return &Random{}
	case "true":
		return &AlwaysTrue{}
	case "false":
		return &AlwaysFalse{}
	default:
		return &Loop{}
	}
}
