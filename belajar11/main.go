package main

import (
	"fmt"
)

func main() {

	/* for index := 0; index < 10; index++ {
		fmt.Println("Perulangan ke", index)
	} */

	/* benar := true
	i := 0

	for benar {
		if i == 3 {
			benar = false
		}
		fmt.Println("Angka ke", i)
		i++
	} */

	/* var index = 0

	for index < 4 {
		fmt.Println("Angka ke", index)
		index++
	} */

	/* var index = 0

	for index < 4 {

		fmt.Println("Angka ke", index)
		index++

		if index == 2 {
			break
		}
	} */

	data := []string{"didik", "Prabowo"}

	for index, v := range data {
		fmt.Println("Perulangan ke", index, v)
	}
}
