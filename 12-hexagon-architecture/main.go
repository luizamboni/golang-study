package main

import (
	"fmt"
	"io"
	"net/http"

	httpserver "go-study.com/m/adapters/http/server"
	"go-study.com/m/adapters/memory/notificator"
	"go-study.com/m/adapters/memory/repo"

	"go-study.com/m/core"
)

func main() {
	logger := core.NewLogger(core.DEBUG)

	ruleRepo := repo.NewRuleRepo()
	// productRepo := repo.NewRuleRepo()
	smsNotificator := notificator.NewSmsNotificator(logger)
	server := httpserver.NewHttpServer(logger)

	classifications := ruleRepo.GetRules()

	categorizer := core.NewCategorizer(logger)

	server.HandleGet("/categorize", func(w http.ResponseWriter, r *http.Request) {

		product, err := server.ParseProductFromReader(r.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("server: could not read request body: %s\n", err))
		}

		category, err := categorizer.CategorizeProductByRules(*product, classifications)
		if err != nil {
			logger.Error(
				fmt.Sprintf("categorize fail by %v", err.Error()),
			)
		}
		if category != "" {
			product.Category = category
			if smsNotificator.Notify(*product) == false {
				logger.Error("fail to notify change")
			}
		}

		io.WriteString(w, fmt.Sprintf("ack! %v \n", product))
	})

	server.HandleGet("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, HTTP!\n")
	})

	server.Init(4000)
}
