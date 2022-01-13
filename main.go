package main

import (
	"fmt"
	) 

func main() {
	// slice
	
	// var gamingConsoles []string
	// gamingConsoles = append(gamingConsoles, "PS 5")
	
	gamingConsoles := []string{"PS5","Nintendo","X BOX"}
	
	
	//fmt.Println(gamingConsoles)

	// iterate slice
	for _, console := range gamingConsoles {
		fmt.Println(console);
	}
}