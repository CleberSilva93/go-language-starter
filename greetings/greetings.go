package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

func Ola(name string) string{
	message := fmt.Sprintf("Olá, %v. Seja bem-vindo!", name)

	return message
}