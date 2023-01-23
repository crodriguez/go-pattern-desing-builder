package main

import (
	"log"

	"github.com/crodriguez/go-pattern-desing-builder/sendmail/domain"
)

func main() {
	var email domain.Email
	email.
		From("crodrigp@gmail.com").
		To([]string{"email1@gmail.com", "email2@gmail.com"}).
		Subject("Este es un email de prueba").
		Body("Este es el cuerpo del correo")

	log.Println(email)
}
