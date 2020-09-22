package main

import (
	"fmt"
)

func cetakData(i interface{}) {
	fmt.Println(i)
}

func main() {
	// // contoh tipe data string
	// st := "Halooo, belajar golang itu mudah lho..."
	// cetakData(st)

	// //tipe slice
	// mp := []int{1, 2, 3, 4, 5}
	// cetakData(mp)

	// var mp interface{}

	// //tipe slice
	// mp = []string{"Satu", "Dua"}

	// var toSlice = strings.Join(mp.([]string), ", ")
	// fmt.Println(toSlice)

	var mp interface{}

	data := map[string]string{
		"Name":    "Didik Prabowo",
		"Address": "Wonogiri",
		"Website": "Kodingin.com",
	}

	mp = map[string]interface{}{
		"data": data,
	}

	fmt.Println(mp)
}
