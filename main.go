package main

import "github.com/rogercoll/terformance/internal/controller"

func main() {
	c := controller.LoadConfig()
	c.Run()
}
