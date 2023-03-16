package main

import (
	"fmt"
	"time"
)

func main() {
	// this not work without tzdata support
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	// time.Local, _ = time.LoadLocation("Europe/Amsterdam")
	time.Local = time.FixedZone("UTC-3", -3*60*60)

	t := time.Now()
	zone_name, offset := t.Zone()
	fmt.Println(t.Format("2006-01-02 15:04:05"), "in timzone", zone_name, offset)
}
