package main

import (
	"fmt"
	) 

func main() {
	// make variable map
	// var myMap map[string]int

	// declare first value of map
	// myMap = map[string]int{}

	// then insert data of map
	// myMap["GO"] = 10
	// myMap["Ruby"] = 9
	// myMap["Laravel"] = 8
	// myMap["PHP"] = 8

	// or make map like below
	// myMap
	// insert value directly on map
	myMap := map[string]string{"GO": "Nice", "Ruby": "Sure", "JS": "Nice Try"}

	for idx, val := range myMap {
		fmt.Println("key: ", idx, " value: ", val)
	}

	fmt.Println("----------------------------------------------------")

	delete(myMap, "Ruby")
	for idx, val := range myMap {
		fmt.Println("key: ", idx, " value: ", val)
	}

	fmt.Println("----------------------------------------------------")

	k, is := myMap["GO"]

	fmt.Println("check if key: " + k +" of map is exist: ", is)

	fmt.Println("map on slice ----------------------------------------------------")

	sliceMap := []map[string]string{
		{"student":"Mario", "value": "A"},
		{"student":"Jun", "value": "B"},
		{"student":"Jey", "value": "C"},
	}
	for _, smap := range sliceMap {
		fmt.Println("murid bernama: ", smap["student"], " mendapat nilai: ", smap["value"])
	}
}