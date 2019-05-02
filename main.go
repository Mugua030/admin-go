package main

import (
	"admin-go/router"
	"log"
)

func main() {
	router := router.Init()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Start server %+v", err)
	}
}
