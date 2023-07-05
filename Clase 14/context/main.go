package main

import (
	"context"
	"fmt"
)

func main(){
	//Background se utiliza para inicializar un contexto
	ctx := context.Background()

	newCtx := addValue(ctx)

	s := newCtx.Value("Prueba")

	fmt.Println(s)
}

	// Atravez de un contexto padre ctx Creamos un nuevo contexto con una clave valor asociada a esta
func addValue(ctx context.Context) context.Context{
	return context.WithValue(ctx, "Prueba", "GO")
}