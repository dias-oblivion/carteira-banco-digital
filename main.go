package main

import (
	"log"

	"github.com/dias-oblivion/PicPay-Simplificado/api"
	"github.com/dias-oblivion/PicPay-Simplificado/database"
)

func main() {
	err := database.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8081")
	server.Start()
}
