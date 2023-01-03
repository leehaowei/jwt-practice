package main

import (
	"fmt"
	"jwt-practice/server/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env") // oad .env file
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from Server")
	api.HandleRequests()
}
