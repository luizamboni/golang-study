package notificator

import (
	"fmt"

	"go-study.com/m/core"
)

type SmsNotificator struct {
	logger *core.Logger
}

func (s SmsNotificator) Notify(product core.Product) bool {
	s.logger.Debug(fmt.Sprintf("send SMS by new category %v", product))
	return true
}

func NewSmsNotificator(logger *core.Logger) *SmsNotificator {

	return &SmsNotificator{
		logger: logger,
	}
}
