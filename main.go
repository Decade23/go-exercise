package main

import (
	"fmt"
	) 

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func(persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}


func main() {
	persegi := Persegi{Sisi: 9}
	fmt.Println("Hitung Luas :" , SeberapaLuas(persegi))
	
	fmt.Println("--------------------------------------")

	persegiPanjang := PersegiPanjang{Panjang: 10, Lebar: 5}
	fmt.Println("Hitung Luas :" , SeberapaLuas(persegiPanjang))
}

func SeberapaLuas(bidangDatar BangunDatar) int {
	return bidangDatar.HitungLuas()
} 