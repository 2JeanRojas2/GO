package main

import (
	"fmt"
)

type Producto struct {
	Nombre string
	Precio float64
	Cantidad int
}
type Servicio struct {
	Nombre string 
	Precio float64
	MinutosTrabajados float64
}
type Mantenimiento struct {
	Nombre string 
	Precio float64
}

func main() {

	productos := []Producto{
        {"Producto 1", 2.5, 3},
        {"Producto 2", 1.99, 5},
        {"Producto 3", 10.0, 1},
    }
	servicios := []Servicio{
        {"Servicio 1", 10.0, 20},
        {"Servicio 2", 20.0, 45},
        {"Servicio 3", 15.0, 60},
    }

	mantenimientos := []Mantenimiento{
		{"Mantenimiento 1", 10.5},
		{"Mantenimiento 2", 15.2},
		{"Mantenimiento 3", 12.5},
	}

	c1 := make(chan float64)
    c2 := make(chan float64)
    c3 := make(chan float64)

    go SumProductos(productos, c1)
    go sumServicios(servicios, c2)
    go sumMantenimiento(mantenimientos, c3)

    
    total := <-c1 + <-c2 + <-c3

    fmt.Printf("Monto total: $%.2f", total)
    
}

func SumProductos(productos []Producto, c chan float64) {
    total := 0.0
    for _, p := range productos {
        total += p.Precio * float64(p.Cantidad)
    }
    c <- total
}

func sumServicios(servicios []Servicio, c chan float64){
	total := 0.0
	for _, s := range servicios {
		duracionEnHoras := s.MinutosTrabajados / 60.0
        if duracionEnHoras < 0.5 {
            duracionEnHoras = 0.5
        }
        total += duracionEnHoras * s.Precio
	}

	c <- total
}

func sumMantenimiento(mantenmientos []Mantenimiento, c chan float64){
	total := 0.0
	for _, m := range mantenmientos{
		total += m.Precio
	}

	c <- total
}
