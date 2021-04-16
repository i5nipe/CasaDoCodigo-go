package main

import "fmt"

func main() {
	var i int

loop:
	for i = 0; i < 10; i++ {
		switch i {
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d\n", i)
			break
		case 5:
			fmt.Println("Quebrando loop, i == 5.")
			break loop
		}
	}
}
