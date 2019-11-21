package service

import (
	"encoding/json"
	"fmt"
)

type Offer struct {
	Name    string
	Price   float64
	Count   float64
	Average float64
}

// run this when json.Unmarshal is call
func (o *Offer) UnmarshalJSON(b []byte) error {

	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})
	fmt.Println("UnmarshalJSON", m)

	if i, ok := m["Average"]; ok == true {
		o.Average = i.(float64)
	}

	if i, ok := m["Count"]; ok == true {
		o.Count = i.(float64)
	}

	if i, ok := m["Price"]; ok == true {
		o.Price = i.(float64)
	}

	// handle name or Name or NAME
	if i, ok := m["name"]; ok == true {
		o.Name = i.(string)
	}

	if i, ok := m["Name"]; ok == true {
		o.Name = i.(string)
	}

	if i, ok := m["NAME"]; ok == true {
		o.Name = i.(string)
	}

	return nil
}
