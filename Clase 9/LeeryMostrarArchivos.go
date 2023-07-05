package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Abrir el archivo customers.txt en modo lectura
	file, err := os.Open("customers.txt")
	if err != nil {
		// Si no se puede abrir el archivo, lanzar un panic
		panic(err)
	}
	defer file.Close()

	// Leer el contenido del archivo
	content, err := ioutil.ReadAll(file)
	if err != nil {
		// Si no se puede leer el archivo, lanzar un panic
		panic(err)
	}

	// Convertir el contenido a string y mostrarlo en pantalla
	fmt.Println(string(content))
}
