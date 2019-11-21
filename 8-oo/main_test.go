package test

import (
	"fmt"
	"testing"

	"example.com/m/service"
)

type StorageMock struct {
	store map[string]string
}

func (s StorageMock) Get(key string) string {
	return "{ \"Name\": \"Mochila\" }"
}

func (s StorageMock) Set(key string, v string, ttl int) error {
	s.store[key] = v
	return nil
}

func TestCacheKeyProduced(t *testing.T) {
	store := map[string]string{"offer-1": "{\"Name\": \"Mochila\"}"}
	serv := service.New(StorageMock{
		store: store,
	})

	got := serv.GetOfferByID("1")

	fmt.Println("got", got)
	if got.Name != "Mochila" {
		t.Errorf("Abs(-1) = %s; want 1", got.Name)
	}
}
