package main

import (
	"api-rest-go/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor..")
	routes.HandleRequest()
}
