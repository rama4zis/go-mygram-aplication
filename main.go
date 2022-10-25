package main

import (
	"github.com/rama4zis/go-mygram-aplication/models"
	"github.com/rama4zis/go-mygram-aplication/router"
)

func main() {
	models.ConnectDatabase()
	r := router.StartApp()

	r.Run()
}
