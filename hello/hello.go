package main

import (
	"fmt"
	"go-starter/greetings"
)

func main(){
	fmt.Println("Hello world :D")

	message := greetings.Hello("Gladys")

	fmt.Println(message)
}