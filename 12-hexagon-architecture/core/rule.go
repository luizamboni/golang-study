package core

type Classification struct {
	Target string
	Rules  []Rule
}

type Rule struct {
	AttributeName  string
	Operation      string
	Operator       string
	ExpectedValues []string
}
