package main

import (
	"fmt"
)

func main() {
	/* var mahasiswa = map[string]string{}

	mahasiswa["Nama"] = "Didik Prabowo"
	mahasiswa["alamat"] = "Wonogiri"

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["alamat"]) */

	/* var mahasiswa = map[string]string{"nama": "Didik Prabowo", "alamat": "Wonogiri"}

	for i, v := range mahasiswa {
		fmt.Println(i, "=", v)
	} */

	/* var mahasiswa = map[string]string{"nama": "Didik Prabowo", "alamat": "Wonogiri"}
	delete(mahasiswa, "alamat")

	for i, v := range mahasiswa {
		fmt.Println(i, "=", v)
	} */

	var mahasiswa = []map[string]string{
		map[string]string{"nama": "Didik Prabowo", "alamat": "Wonogiri"},
		map[string]string{"nama": "Rudi Subgyo", "alamat": "Solo"},
	}

	for i, v := range mahasiswa {
		fmt.Println(i, "= nama :", v["nama"], "alamat:", v["alamat"])
	}
}
