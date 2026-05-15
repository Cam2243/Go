package main

import "fmt"

func main() {
	var nombre string = "Gopher"
	var edad int = 25

	puntos := 100
	activo := true

	fmt.Printf("Jugador: %s, Edad: %d\n", nombre, edad)
	fmt.Printf("Puntos: %d, Activo: %v\n", puntos, activo)
}