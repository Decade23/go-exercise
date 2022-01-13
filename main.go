package main

import (
	"strconv"
	"fmt"
	) 

func main() {
	var name string = "Dedi Fardiyanto"
	age := 23
	fmt.Println("Hello World " + name + " Age: " + strconv.Itoa(age)) 
	
	// loop
	// for normal
	// for i := 1; i <= 100; i ++ {
	// 	fmt.Println("im learn of go : ", i)
	// 	if i == 100 {
	// 		fmt.Println("------------------------------------------")
	// 	}
	// }

	// for style while
	// a := 1
	// for a <= 100 {
	// 	fmt.Println("learn of go : ", a)
	// 	a++
	// }

	// for style ForEach
	// sentence := "here we are learning of GO"
	// for i, l := range sentence {
	// 	fmt.Println("index: ", i, " letter: ", string(l))
	// }

	// for i, l := range sentence {
		// ls := string(l)
		// using modulus
		// if i % 2 != 0 {
		// 	fmt.Println("index : ", i, " letter: ", ls)
		// }

		// filter print by text vocal
		// switch ls {
		// case "a", "i", "u", "e", "o":
		// 	fmt.Println("index : ", i, " letter: ", ls)
		// }
	// }

	// array
	arr := [...]string {
		"c", "c#", "go", "php", "javascript",
	}
	fmt.Println("array length : ", len(arr))
	for idx, let := range arr {
		fmt.Println("idx : ", idx, " word: ", let)
	}

}