package api

import (
	"fmt"
	"io/ioutil"
	"jwt-practice/client/token"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := token.GenerateJWT()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	client := &http.Client{}
	requestURL := os.Getenv("SERVER") + os.Getenv("SERVER_PORT") + "/"
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)

	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprint(w, string(body))
}

func HandleRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("CLIENT_PORT"), nil))

}
