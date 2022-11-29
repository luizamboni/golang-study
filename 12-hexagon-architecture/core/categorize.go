package core

import (
	"fmt"
	"strings"
)

func getAttributeValue(attrName string, product Product) (string, error) {
	if attrName == "name" {
		return product.Name, nil
	}

	if attrName == "ean" {
		return product.Ean, nil
	}

	return "", fmt.Errorf("unknow field")
}

type Comparator func(string, string) bool

func Equals(value string, expected string) bool {
	return value == expected
}

func Contains(value string, expected string) bool {
	return strings.Contains(value, expected)
}

func MatchAny(comp Comparator, value string, expecteds []string) bool {
	for _, expected := range expecteds {
		matched := comp(value, expected)
		if matched == true {
			return matched
		}
	}
	return false
}

func MatchAll(comp Comparator, value string, expecteds []string) bool {
	for _, expected := range expecteds {
		matched := comp(value, expected)
		if matched != true {
			return matched
		}
	}
	return true
}

func CategorizeProductByRules(product Product, classifications []Classification) (string, error) {
	for _, classification := range classifications {
		matchedAll := false
		for _, rule := range classification.Rules {
			attrValue, err := getAttributeValue(rule.AttributeName, product)
			if err != nil {
				return "", err
			}

			matched := false

			if rule.Operator == "equal" {
				if rule.Operation == "OR" {
					matched = MatchAny(Equals, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = MatchAll(Equals, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "contains" {
				if rule.Operation == "OR" {
					matched = MatchAny(Contains, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = MatchAll(Contains, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "startWith" {
				if rule.Operation == "OR" {
					matched = MatchAny(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = MatchAll(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "notStartWith" {
				if rule.Operation == "OR" {
					matched = !MatchAny(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = !MatchAll(strings.HasPrefix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "endWith" {
				if rule.Operation == "OR" {
					matched = MatchAny(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = MatchAll(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
			}

			if rule.Operator == "notEndWith" {
				if rule.Operation == "OR" {
					matched = !MatchAny(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
				if rule.Operation == "AND" {
					matched = !MatchAll(strings.HasSuffix, attrValue, rule.ExpectedValues)
				}
			}

			fmt.Printf("field '%s' %s %s %s against '%s': %v \n",
				rule.AttributeName, rule.Operator, rule.Operation,
				rule.ExpectedValues, attrValue,
				matched,
			)

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
