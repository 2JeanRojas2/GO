package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	/*
		Verbo utilizado: GET, POST, PUT, etc.
		Fecha y hora: podemos utilizar el paquete time.
		URL de la consulta: localhost:8080/products
		Tamaño en bytes: tamaño de la consulta.
		Elapsed Time
	*/
	return func(c *gin.Context) {
		t := time.Now()
		verb := c.Request.Method
		path := c.Request.RequestURI

		// Paso el control al otro middleware/handler
		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}
		elapsed := time.Since(t)
		fmt.Printf("time: %v\npath: %s\nverb: %s\nsize: %d\nelapsed time: %v", t, path, verb, size, elapsed)
	}
}
