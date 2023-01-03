package api

import (
	"fmt"
	"jwt-practice/server/middleware"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func HandleRequests() {
	http.Handle("/", middleware.IsAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}
