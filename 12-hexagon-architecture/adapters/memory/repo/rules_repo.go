package repo

import (
	"go-study.com/m/core"
)

type RulesRepo struct {
}

func (r RulesRepo) GetRules() []core.Classification {
	classifications := []core.Classification{
		{
			Target: "coisa veloz",
			Rules: []core.Rule{
				{
					AttributeName:  "name",
					Operation:      "AND",
					Operator:       "endWith",
					ExpectedValues: []string{"veloz"},
				},
				{
					AttributeName:  "name",
					Operation:      "AND",
					Operator:       "notStartWith",
					ExpectedValues: []string{"carro"},
				},
			},
		},
		{
			Target: "Carro veloz",
			Rules: []core.Rule{
				{
					AttributeName:  "name",
					Operation:      "AND",
					Operator:       "contains",
					ExpectedValues: []string{"carro", "veloz"},
				},
			},
		},
	}
	return classifications
}

// or panic
func NewRuleRepo() *RulesRepo {

	return &RulesRepo{}
}
