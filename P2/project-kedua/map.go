package main

import "log"

func main() {

	// map with iterface
	car := map[string]interface{}{
		"type":              "Sedan",
		"registration_year": 2022,
	}
	log.Println(car["type"])


	// map & slice harus kuat
}