package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds db.AuthoData
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println("error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, found := GetUserDB(creds.Phone)
	fmt.Println("cred ", creds.Opt, " data", users.Opt)
	if found {
		if creds.Opt == users.Opt {
			fmt.Println("correct cred")
			w.WriteHeader(http.StatusAccepted)
			return
		} else {
			fmt.Println("incorrect cred")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		fmt.Println("not found")
		var jwtKey = []byte("my_secret_key")
		user := db.User{
			Id:    10,
			Name:  "waleed",
			Phone: creds.Phone,
			Opt:   CreateOPT(),
			Token: "5456546",
		}
		InsertUserDB(user)

		expirationTime := time.Now().Add(5 * time.Minute)
		claims := db.Claims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}

}
