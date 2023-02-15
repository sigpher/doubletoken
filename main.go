package main

import "doubletoken/router"

func main() {
	router := router.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		panic("Running server failed, err")
	}
}
