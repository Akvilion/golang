package main

import (
	"fmt"
	"xxx/geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
