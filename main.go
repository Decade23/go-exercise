package main

import "fmt"

type Sentence struct {
	NumberA   int
	Sentences []string
}

func (sentence Sentence) display() {
	fmt.Printf("No: %d", sentence.NumberA)
	for _, v := range sentence.Sentences {
		fmt.Printf("panjang %s karakter dari kalimat ini: ", v)
	}
}

func main() {
	numberA := 10
	numberB := &numberA

	// change value of numberB shall effect to  numberA because its change the value of computerized
	*numberB = 20
	fmt.Println("number A: ", numberA)
	fmt.Println("---------------")
	fmt.Println("number B: ", *numberB)

	fmt.Println("---------------")

}
