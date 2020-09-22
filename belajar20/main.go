package main

import (
	"fmt"
)

// type Pekerja struct {
// 	Nama string
// 	Gaji int
// }

// func (p Pekerja) lihatPekerja() {
// 	fmt.Println("Nama\t :", p.Nama)
// 	fmt.Println("Gaji\t :", p.Gaji)
// }

// func main() {
// 	p1 := Pekerja{
// 		Nama: "Charly Van Houten",
// 		Gaji: 1000000,
// 	}

// 	p1.lihatPekerja()
// }

// const UMR int = 2000000

// type Pekerja struct {
// 	Nama string
// 	Gaji int
// }

// func (p Pekerja) lihatPekerja() bool {
// 	fmt.Println("Nama\t :", p.Nama)
// 	fmt.Println("Gaji\t :", p.Gaji)
// 	if p.Gaji < UMR {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func main() {
// 	p1 := Pekerja{
// 		Nama: "Charly Van Houten",
// 		Gaji: 1000000,
// 	}

// 	fmt.Println(p1.lihatPekerja())
// }

// type (
// 	Category struct {
// 		Name string
// 	}

// 	Post struct {
// 		Title string
// 	}
// )

// func (c Category) lihatData() {
// 	fmt.Println(c)
// }
// func (p Post) lihatData() {
// 	fmt.Println(p)
// }

// func main() {
// 	fmt.Printf("From Category\n")
// 	cats := Category{
// 		Name: "Berita",
// 	}

// 	cats.lihatData()

// 	fmt.Printf("From Post\n")

// 	p := Post{
// 		Title: "Belajar Golang",
// 	}
// 	p.lihatData()
// }

type (
	Category struct {
		Name string
	}
)

// func (c Category) lihatKategori() {
// 	c.Name = "Berita"
// }
// func (c Category) lihatKategoriLagi() {
// 	c.Name = "Teknologi"
// }

// func main() {
// 	cats := Category{
// 		Name: "Info",
// 	}

// 	fmt.Println(cats)

// 	(&cats).lihatKategori()
// 	fmt.Println(cats)

// 	(&cats).lihatKategoriLagi()
// 	fmt.Println(cats)
// }

type val int

func (i val) jumlah() val {
	s := i * val(2)
	return s
}

func main() {
	numb := val(10)
	fmt.Println(numb)
	fmt.Println(numb.jumlah())

}
