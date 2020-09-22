package main

import (
	"fmt"
)

// Profil
// type Profil struct {
// 	ID      int
// 	Name    string
// 	Age     int
// 	Address string
// }

// func main() {
// 	var profil Profil

// 	profil.ID = 1
// 	profil.Name = "Didik Prabowo"
// 	profil.Age = 20
// 	profil.Address = "Wuryantoto, Wonogiri"

// 	fmt.Println(profil)
// 	fmt.Println("===========================")
// 	fmt.Printf("Nama \t: %v\n", profil.Name)
// 	fmt.Printf("Umur \t: %d\n", profil.Age)
// 	fmt.Printf("Alamat \t: %v\n", profil.Address)
// 	fmt.Println("===========================")
// }

// type Profil struct {
// 	ID      int
// 	Name    string
// 	Age     int
// 	Address string
// }

// func main() {

// 	profil := Profil{
// 		ID:      1,
// 		Name:    "Didik Prabowo",
// 		Age:     40,
// 		Address: "Wuryantoto, Wonogiri",
// 	}

// 	fmt.Println(profil)
// 	fmt.Println("===========================")
// 	fmt.Printf("Nama \t: %v\n", profil.Name)
// 	fmt.Printf("Umur \t: %d\n", profil.Age)
// 	fmt.Printf("Alamat \t: %v", profil.Address)
// }

// Profil
type (
	Education struct {
		SMA string
		SMP string
	}

	Profil struct {
		ID        int
		Name      string
		Age       int
		Address   string
		Education Education
	}
)

func main() {

	profil := Profil{
		ID:      1,
		Name:    "Didik Prabowo",
		Age:     40,
		Address: "Wuryantoto, Wonogiri",
		Education: Education{
			SMA: "SMA N 1 Wuryantoro",
			SMP: "SMP N 1 Wuryantoro",
		},
	}

	fmt.Println(profil)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", profil.Name)
	fmt.Printf("Umur \t: %d\n", profil.Age)
	fmt.Printf("Alamat \t: %v", profil.Address)
	fmt.Printf("Sekolah \t: %v", profil.Education)
}
