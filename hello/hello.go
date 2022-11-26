package main

import (
	"fmt"
	"go-starter/greetings"
)

func main(){
	fmt.Println("Hello world :D")

	message := greetings.Hello("Gladys")

	messagept := greetings.Ola("Gladys")

	fmt.Println(message)
	fmt.Println(messagept)
}