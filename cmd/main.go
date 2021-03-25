package main

import (
	"log"
	// unittest "unit-test"
	"unit-test/customer"
	_ "unit-test/docs"
)


func main() {
	// unittest.Test()
	log.Panic(customer.StartServer())
}
