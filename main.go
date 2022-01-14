package main

import (
	"errors"
	"fmt"
	) 

func main() {
	// luas, kel := calcPref(20, 4)
	// fmt.Println("luas: ", luas, " keliling: ", kel)

	quiz()
}

func calc(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * ( panjang + lebar)

	return luas, keliling
}

func calcPref(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * ( panjang + lebar)
	return
}

func quiz() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("quiz scores: ", total)

	res, err := calQuiz(10, 2, "/")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	// res, err := calQuiz(10, 2, "+")
	// res, err := calQuiz(10, 2, "-")
	// res, err := calQuiz(10, 2, "*")
	// res, err := calQuiz(10, 2, "/")
}

func calQuiz(number1, number2 int, operation string) (res int, err error) {
	switch operation {
		case "+":
			res = number1 + number2
			case "-":
				res = number1 - number2
				case "*":
					res = number1 * number2
					case "/":
						res = number1 / number2
						default:
							err = errors.New("Unkwon Operation")
	}
	return
}

func sum(scores []int) (res int) {
	
	for _, score := range scores {
		res += score
	}
	return
}