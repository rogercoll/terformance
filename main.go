package main

import (
	"log"
	"os"

	"github.com/rogercoll/terformance/internal/controller"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Provide filename to create list")
	}
	c, err := controller.LoadConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c.Run()
}
