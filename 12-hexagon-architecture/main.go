package main

import (
	"fmt"

	"go-study.com/m/adapters/memory/notificator"
	"go-study.com/m/adapters/memory/repo"
	"go-study.com/m/core"
)

func main() {

	ruleRepo := repo.NewRuleRepo()
	productRepo := repo.NewRuleRepo()
	smsNotificator := notificator.NewSmsNotificator()

	classifications := ruleRepo.GetRules()

	for _, product := range productRepo.GetProducts() {

		category, err := core.CategorizeProductByRules(product, classifications)
		if err != nil {
			fmt.Println("categorize fail by", err.Error())
		}
		if category != "" {
			product.Category = category
			if smsNotificator.Notify(product) == false {
				fmt.Println("fail to notify change")
			}
		}
	}
}
