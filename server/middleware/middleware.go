package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// var mySigningKey = []byte("mysupersecretphrase")
var mySigningKey = []byte(os.Getenv("MY_SIGNING_KEY"))

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized") // no token
		}
	})
}
