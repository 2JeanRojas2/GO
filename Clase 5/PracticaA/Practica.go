package main

import "fmt"

type Usuarios struct{
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contraseña string
}

func (u *Usuarios) InformacionUsuario(){
	fmt.Println("Nombre:", u.Nombre," Edad:", u.Edad," Correo:", u.Correo," Contraseña:", u.Contraseña)
}

func (u *Usuarios) CambiarNombre(nombre string) {
	u.Nombre = nombre
}

func (u *Usuarios) CambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuarios) CambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuarios) CambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main(){
	nuevo := Usuarios{
		Nombre: "Jean",
		Apellido: "Rojas",
		Edad: 24,
		Correo: "jean@gmail.com",
		Contraseña: "12345",
	}	

	nuevo.InformacionUsuario()
	nuevo.CambiarNombre("Luis")
	nuevo.CambiarEdad(21)
	nuevo.CambiarCorreo("Luis@gmail.com")
	nuevo.CambiarContraseña("45667")
	nuevo.InformacionUsuario()
}