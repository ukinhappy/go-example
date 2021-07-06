package main

import (
	"fmt"
	"github.com/uber/h3-go/v3"
)

func main() {
	geo := h3.GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	fmt.Println(h3.FromGeo(geo, 8))
}
