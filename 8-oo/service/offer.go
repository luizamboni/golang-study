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

	o.Average = m["Average"].(float64)
	o.Count = m["Count"].(float64)

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

	// o.Name = m["name"].(string)
	o.Price = m["Price"].(float64)

	return nil
}
