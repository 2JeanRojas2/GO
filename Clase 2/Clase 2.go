package main

import (
	"fmt"
)

func clase2() {

	const (
		Suma = "+"
		Resta = "-"
		Multiplicar = "*"
		Dividir = "/"
	)

	func opSuma(valor1, valor2 float64) float64 {
		return valor1 + valor2
	}

	func opResta(valor1, valor2 float64) float64{
		return valor1 - valor2
	}

	func opMultiplicar(valor1, valor2 float64) float64{
		return valor1 * valor2
	}

	func opDividir(valor1, valor2 float64) float64{
		if valor2 == 0 {
			return 0
		}
		return valor1 / valor2
	}

	func operacionAritmetica(operador string, valores ...float64 ) float64{
		switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case Resta:
		return orquestadorOperaciones(valores, opResta)
	case Multiplicar:
		return orquestadorOperaciones(valores, opMultiplicar)
	case Dividir:
		return orquestadorOperaciones(valores, opDividir)
	}
	return 0;
	}

	func orquestadorOperaciones(valores []float64, 
		operacion func (valor1, valor2 float64) float64)float64{
			var resultado float64
			for i, valor: range valores{
				if i == 0 {
				resultado = valor
				}else{
					resultado = operacion(resultado, valor)
				}
			}
			return resultado
		}

	fmt.Println(operacionAritmetica(Suma, 2 3 4 5 6 7 8 6 5 4 3))
}