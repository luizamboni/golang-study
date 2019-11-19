package service

import (
	"fmt"

	"example.com/m/interfaces"
)

type Service struct {
	c interfaces.Storage
}

func (s Service) GetOfferByID(id string) (string, error) {
	fmt.Println(s, s.c)
	return s.c.Get("offer-" + id)
}

func New(c interfaces.Storage) *Service {
	return &Service{c}
}
