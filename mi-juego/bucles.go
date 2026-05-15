package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	contador := 0
	for contador < 3 {
		fmt.Println("Contando:", contador)
		contador++
	}

	for {
		fmt.Println("Bucle infinito")
		break 
	}
}