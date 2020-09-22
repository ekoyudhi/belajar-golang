package main

import "fmt"

// type Kendaraan interface {
// 	GetWarna()
// }

// type Mobil struct {
// 	Depan    string
// 	Belakang string
// 	Atas     string
// 	Dalam    string
// }

// type Motor struct {
// 	Depan    string
// 	Belakang string
// }

// func (m Mobil) GetWarna() {
// 	fmt.Println(m)
// }

// func CetakWarna(w Kendaraan) {
// 	fmt.Println(w)
// }

// // func (m Motor) GetWarna() {
// // 	fmt.Println(m)
// // }

// func main() {
// 	var k Kendaraan
// 	mobil := Mobil{
// 		Depan:    "Biru",
// 		Belakang: "Biru",
// 		Atas:     "Pink",
// 		Dalam:    "Putih",
// 	}

// 	k = mobil
// 	k.GetWarna()
// 	CetakWarna(k)

// 	// motor := Motor{
// 	// 	Depan:    "Green",
// 	// 	Belakang: "Black",
// 	// }

// 	// k = motor
// 	// k.GetWarna()
// }

type WarnaModif interface {
	GetWarnaModif()
}

type WarnaOriginal interface {
	GetWarnaOriginal()
}

type Warna interface {
	WarnaModif
	WarnaOriginal
}

type Mobil struct {
	Depan    string
	Belakang string
	Atas     string
	Dalam    string
}

func (m Mobil) GetWarnaModif() {
	fmt.Println(m)
}

func (m Mobil) GetWarnaOriginal() {
	fmt.Println(m)
}

func main() {
	var k Warna
	mobil := Mobil{
		Depan:    "Biru",
		Belakang: "Biru",
		Atas:     "Pink",
		Dalam:    "Putih",
	}

	k = mobil
	k.GetWarnaModif()
}
