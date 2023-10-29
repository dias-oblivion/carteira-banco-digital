package main

import (
	"fmt"
	"log"

	"github.com/dias-oblivion/PicPay-Simplificado/api"
	"github.com/dias-oblivion/PicPay-Simplificado/database"
)

func main() {
	storage, err := database.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", storage)

	server := api.NewAPIServer(":8081", storage)
	server.Start()
}
