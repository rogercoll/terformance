package main

import "github.com/rogercoll/terformance"

func main() {
	c := terformance.LoadConfig()
	c.Run()
}
