package main

import (
	"fiber-swagger/controller"
	_ "fiber-swagger/docs"
)

func main() {
	controller.StartServer()
}
