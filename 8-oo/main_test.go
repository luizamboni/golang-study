package test

import (
	"testing"

	"example.com/m/service"
)

type StorageMock struct {
	store map[string]string
}

func (s StorageMock) Get(key string) string {
	return s.store[key]
}

func (s StorageMock) Set(key string, v string, ttl int) error {
	s.store[key] = v
	return nil
}

func TestCacheKeyProduced(t *testing.T) {
	store := map[string]string{
		"offer-1": "{\"Name\": \"Mochila\"}",
		"offer-2": "{\"name\": \"Mochila\"}",
		"offer-3": "{\"NAME\": \"Mochila\"}",
	}

	serv := service.New(StorageMock{
		store: store,
	})

	examples := []struct {
		id       string
		expected string
	}{
		{"1", "Mochila"},
		{"2", "Mochila"},
		{"3", "Mochila"},
	}

	for _, example := range examples {
		got := serv.GetOfferByID(example.id)

		if got.Name != example.expected {
			t.Errorf("Mochila = %s; want \"Mochila\"", got.Name)
		}
	}
}
