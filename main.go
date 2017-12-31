package main

import "github.com/alberii/akeome/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
