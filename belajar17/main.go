package main

import (
	"fmt"
)

// func main() {

// 	v := func(nilai int) int {
// 		return nilai
// 	}

// 	result := checkNilai(v)
// 	fmt.Println(result)
// }

// func checkNilai(nilai func(int) int) int {
// 	return nilai(10)
// }

func main() {

	v := func(nilai int) int {
		return nilai
	}

	name, _ := checkNilai("didikprabowo", v)
	fmt.Println(name)
}

func checkNilai(name string, nilai func(int) int) (string, int) {
	return name, nilai(100)
}
