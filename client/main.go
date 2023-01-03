package main

import (
	"fmt"
	"jwt-practice/client/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env") // oad .env file
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from client")

	api.HandleRequest()
}
