package main

import (
	"fmt"
)

func main() {
	/* var name [5]string

	name[0] = "Didik Prabowo"
	name[1] = "Prabowo Didik"
	name[3] = "Mas Didik Prabowo"
	name[4] = "Prabowo Golang"

	fmt.Println(name[0], name[1], name[2], name[3], name[4]) */

	/* ar name [5]string

	name[0] = "Didik Prabowo"
	name[1] = "Prabowo Didik"
	name[3] = "Mas Didik Prabowo"
	name[4] = "Prabowo Golang"

	for i, v := range name {
		fmt.Println("Index ke ", i, "=>", v)
	} */

	/* var name [5]string

	name[0] = "Didik Prabowo"
	name[1] = "Prabowo Didik"
	name[3] = "Mas Didik Prabowo"
	name[4] = "Prabowo Golang"

	name[len(name)-1] = ""

	for index := 0; index < len(name); index++ {
		fmt.Println("Index ke ", index, "=>", name[index])
	} */

	fmt.Println("Array Multi Dimensi")
	var angka = [5][2]int{{0, 0}, {2, 2}, {3, 4}, {5, 6}, {7, 8}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Data ke [%d][%d] = %d\n", i, j, angka[i][j])
		}
	}
}
