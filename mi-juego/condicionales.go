package main

import "fmt"

func main() {
	puntos := 85
	if puntos >= 90 {
		fmt.Println("Excelente")
	} else if puntos >= 70 {
		fmt.Println("Bien")
	} else {
		fmt.Println("Sigue practicando")
	}
}