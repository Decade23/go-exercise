package main

import (
	"fmt"
	) 

func main() {
	// hitung rata-rata
	// scores := [8]int{100, 80, 97, 85, 90, 65, 78, 84}
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var totalScore int
	// for _, score := range scores {
	// 	totalScore += score
	// }
	// // average of scores is
	// // fmt.Println("Total ", float64(totalScore))
	// fmt.Println("Rata-rata dari score tsb @ ", float64(totalScore) / float64(len(scores)))

		scores := [8]int{100, 80, 97, 85, 90, 65, 78, 84}
		var goodScores []int

		for _, score := range scores {
			if score >= 90 {
				goodScores = append(goodScores, score)
			}
		}
		fmt.Println("Good scores is ", goodScores)
}