package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
1. Crear una estructura Usuario y otra Post que representen los datos que encontrarán dentro de cada archivo JSON.
2. Leer los archivos users.json y posts.json:
	a. Crear un método GET que permite imprimir la lista completa de usuarios en formato JSON.
	b. Crear un método GET que permite recuperar los datos de un único usuario, en formato JSON, recibiendo el ID del mismo por parámetro en la URL.
	c. Crear un método GET que permite recuperar todos los posts realizados por un usuario en particular. Imprimirlo en formato JSON.
	d. Crear una ruta /search que permite buscar por parámetro los usuarios cuyo nombre contenga lo ingresado en el parámetro name (total o parcialmente).

*/

type Geo struct {
	Lat string `json:"lat"`
	Lng string  `json:"lng"`
}

type Address struct {
	Street string  `json:"street"`
	Suite string  `json:"suite"`
	City string  `json:"city"`
	Zipcode string  `json:"zipcode"`
	Geo Geo  `json:"geo"`
}

type Company struct {
	Name string  `json:"name"`
	CatchPhrase string  `json:"catchPhrase"`
	Bs string  `json:"bs"`
}

type Usuario struct {
	Id int  `json:"id"`
	Name string  `json:"name"`
	Username string  `json:"username"`
	Email string  `json:"email"`
	Address Address  `json:"address"`
	Phone string  `json:"phone"`
	Website string  `json:"website"`
	Company Company  `json:"company"`
}

type Post struct {
	UserId int  `json:"userId"`
	Id int  `json:"id"`
	Title string  `json:"title"`
	Body string  `json:"body"`
}

func initializeFiles() ([]Post, []Usuario) {
	posts, err1 := os.ReadFile("posts.json")
	if err1 != nil {
		fmt.Println("Falló el file de posts")
	}
	users, err2 := os.ReadFile("users.json")
	if err2 != nil {
		fmt.Println("Falló el file de users")
	}

	var Posts []Post
	var Users []Usuario
	json.Unmarshal(posts, &Posts)
	json.Unmarshal(users, &Users)

	return Posts, Users
}

func main() {
	router := gin.Default()

	variable := router.Group("/users")
	variable.GET("", imprimirTodos())
	variable.GET("/:id", buscarUser())
	variable.GET("/posts/:id", postsDeUsuario())
	variable.GET("/search", buscarPorNombre())

	router.Run()
}

func imprimirTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, users := initializeFiles()
		ctx.JSON(200, users)
		return
	}
}

func buscarUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, users := initializeFiles()
		id := ctx.Param("id")
		idInt, _ := strconv.ParseInt(id, 10, 64)
		for _, v := range users {
			if int64(v.Id) == idInt {
				ctx.JSON(200, v)
				return
			}
		}
		ctx.JSON(404, "Mal")
		return
	}
}

func postsDeUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, _ := initializeFiles()
		id := ctx.Param("id")
		idInt, _ := strconv.ParseInt(id, 10, 64)
		var userPosts []Post
		for _, v := range posts {
			if int64(v.UserId) == idInt {
				userPosts = append(userPosts, v)
			}
		}
		ctx.JSON(200, userPosts)
		return
	}
}

func buscarPorNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, users := initializeFiles()
		name := ctx.Query("name")
		var usersRes []Usuario

		for _, v := range users {
			if strings.Contains(v.Name, name) {
				usersRes = append(usersRes, v)
			}
		}
		ctx.JSON(200, usersRes)
		return
	}
}