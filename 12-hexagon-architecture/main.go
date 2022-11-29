package main

import (
	"fmt"

	memory "go-study.com/m/adapters/memory"
	core "go-study.com/m/core"
)

func main() {

	ruleRepo := memory.NewRuleRepo()
	productRepo := memory.NewRuleRepo()

	classifications := ruleRepo.GetRules()

	for _, product := range productRepo.GetProducts() {

		category, err := core.CategorizeProductByRules(product, classifications)
		if err != nil {
			fmt.Println("categorize fail by", err.Error())
		}
		fmt.Println("category:", category)
	}
}
