package core

import (
	"fmt"
	"strings"
)

type Categorizer struct {
	logger *Logger
}

func (s Categorizer) getAttributeValue(attrName string, product Product) (string, error) {
	if attrName == "name" {
		return product.Name, nil
	}

	if attrName == "ean" {
		return product.Ean, nil
	}

	return "", fmt.Errorf("unknow field")
}

type Comparator func(string, string) bool

func (s Categorizer) Equals(value string, expected string) bool {
	return value == expected
}

func (s Categorizer) Contains(value string, expected string) bool {
	return strings.Contains(value, expected)
}

func (s Categorizer) MatchAny(comp Comparator, value string, expecteds []string) bool {
	for _, expected := range expecteds {
		matched := comp(value, expected)
		if matched == true {
			return matched
		}
	}
	return false
}

func (s Categorizer) MatchAll(comp Comparator, value string, expecteds []string) bool {
	for _, expected := range expecteds {
		matched := comp(value, expected)
		if matched != true {
			return matched
		}
	}
	return true
}

func (s Categorizer) CategorizeProductByRules(product Product, classifications []Classification) (string, error) {
	for _, classification := range classifications {
		matchedAll := false
		for _, rule := range classification.Rules {
			attrValue, err := s.getAttributeValue(rule.AttributeName, product)
			if err != nil {
				return "", err
			}

			matched := false

			if rule.Operator == "equal" {
				if rule.Operation == "OR" {
					matched = s.MatchAny(s.Equals, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = s.MatchAll(s.Equals, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "contains" {
				if rule.Operation == "OR" {
					matched = s.MatchAny(s.Contains, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = s.MatchAll(s.Contains, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "startWith" {
				if rule.Operation == "OR" {
					matched = s.MatchAny(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = s.MatchAll(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "notStartWith" {
				if rule.Operation == "OR" {
					matched = !s.MatchAny(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = !s.MatchAll(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "endWith" {
				if rule.Operation == "OR" {
					matched = s.MatchAny(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = s.MatchAll(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "notEndWith" {
				if rule.Operation == "OR" {
					matched = !s.MatchAny(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = !s.MatchAll(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
			}

			s.logger.Info(fmt.Sprintf("field '%s' %s %s %s against '%s': %v \n",
				rule.AttributeName, rule.Operator, rule.Operation,
				rule.ExpectedValues, attrValue,
				matched,
			))

			if matched == false {
				matchedAll = false
				break
			}

			matchedAll = true
		}
		if matchedAll == true {
			return classification.Target, nil
		}
	}

	return "", nil

}

func NewCategorizer(logger *Logger) *Categorizer {
	return &Categorizer{
		logger: logger,
	}
}
