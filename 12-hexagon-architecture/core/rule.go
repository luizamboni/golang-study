package core

import "fmt"

type Classification struct {
	Target string
	Rules  []Rule
}

func (s Classification) String() string {
	return fmt.Sprintf(
		"<Target: '%v' Rules: '%v' />",
		s.Target, s.Rules,
	)
}

type Rule struct {
	AttributeName  string
	Operation      string
	Operator       string
	ExpectedValues []string
}

func (s Rule) String() string {
	return fmt.Sprintf(
		"<AttributeName: '%v' Operation: '%v' Operator: '%v' ExpectedValues: '%v' />",
		s.AttributeName, s.Operation, s.Operator, s.ExpectedValues,
	)
}
