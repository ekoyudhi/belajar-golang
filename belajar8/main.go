package main

import "fmt"

func main() {
	var nama, mataKuliah string
	var nilai, nim int
	fmt.Println("Masukkan NIM Anda : ")
	fmt.Scan(&nim)
	fmt.Println("Masukkan Nama Anda: ")
	fmt.Scanf("%s", &nama)
	fmt.Println("Masukkan Mata Kuliah Anda: ")
	fmt.Scanf("%s", &mataKuliah)
	fmt.Println("Masukkan Nilai Anda: ")
	fmt.Scanln(&nilai)

	fmt.Printf("Nama(NIM) => %s %d \n", nama, nim)
	fmt.Printf("Matkul(Nilai) => %s %d \n", mataKuliah, nilai)
}
