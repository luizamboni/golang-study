package notificator

import (
	"fmt"

	"go-study.com/m/core"
)

type SmsNotificator struct {
}

func (s SmsNotificator) Notify(product core.Product) bool {
	fmt.Println("send SMS by new category", product)
	return true
}

func NewSmsNotificator() *SmsNotificator {

	return &SmsNotificator{}
}
