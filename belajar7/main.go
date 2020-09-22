package main

import "fmt"

func main() {
	var nama, alamat, sekolah string
	var umur int

	nama = "Eko Yudhi Prastowo"
	alamat = "Kudus"
	sekolah = "Smada"

	umur = 21
	tinggibadan := 172
	beratbadan := 80

	fmt.Println("========Biodata=========")
	fmt.Println("Nama : " + nama)
	fmt.Println("Alamat : " + alamat)
	fmt.Println("Sekolah : " + sekolah)
	fmt.Println("Umur : ", umur)
	fmt.Println("Tinggi Badan : ", tinggibadan)
	fmt.Println("Berat Badan : ", beratbadan)

}
