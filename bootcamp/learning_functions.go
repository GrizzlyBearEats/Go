package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "LA", "Palo Alto", "Burlingame":
		region, continent = "California", "North America"
	default:
		region, continent = "Unknown", "Uknown"
	}
	return region, continent
}

func main() {
	region, continent := location("LA")
	fmt.Printf("Matt lives in %s, %s", region, continent)
}
