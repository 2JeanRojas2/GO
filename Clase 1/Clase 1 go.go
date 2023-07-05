package main

import "fmt"

func main() {

	var precio float32 = 100;
	var porcentaje float32 = 10;

	var descuento = precio * porcentaje / 100;
	
	fmt.Println(descuento);
}