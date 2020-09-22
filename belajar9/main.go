package main

import "fmt"

func main() {
	//Set Nilai Awal
	angka1 := 10
	angka2 := 5

	//Operator Penjumlahan
	hasiljumlah := angka1 + angka2
	fmt.Println(fmt.Sprintf("Hasil penjumlahan dari %d dan %d adalah %d Bos..", angka1, angka2, hasiljumlah))

	//Operator Pengurangan
	hasilkurang := angka1 - angka2
	fmt.Println(fmt.Sprintf("Hasil Pengurangan dari %d dan %d adalah %d Bos..", angka1, angka2, hasilkurang))

	//Operator Perkalian
	hasilkali := angka1 * angka2
	fmt.Println(fmt.Sprintf("Hasil Perkalian dari %d dan %d adalah %d Bos..", angka1, angka2, hasilkali))

	//Operator Pembagian
	hasilbagi := angka1 / angka2
	fmt.Println(fmt.Sprintf("Hasil Pembagian dari %d dan %d adalah %d Bos..", angka1, angka2, hasilbagi))

	//Operator Perkalian
	hasilsisa := angka1 % angka2
	fmt.Println(fmt.Sprintf("Hasil Sisa Hasil Bagi dari %d dan %d adalah %d Bos..", angka1, angka2, hasilsisa))

}
