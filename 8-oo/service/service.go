package service

import (
	"encoding/json"
	"fmt"

	"example.com/m/interfaces"
)

type Service struct {
	c interfaces.Storage
}

func (s Service) GetOfferByID(id string) *Offer {
	fmt.Println(s, s.c)
	v := s.c.Get("offer-" + id)

	offer := new(Offer)
	json.Unmarshal([]byte(v), offer)
	return offer
}

func New(c interfaces.Storage) *Service {
	return &Service{c}
}
