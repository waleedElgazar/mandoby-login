package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds db.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println("error")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("there is error happened\n try again"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	phone := creds.Phone
	users, found := GetUserAutho(phone)
	if found {
		if creds.Opt == users.Opt {
			_, founded := GetUserDB(phone)
			if founded {
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("accepted"))
				w.WriteHeader(http.StatusAccepted)
				return
			} else {
				user := db.User{
					Name:  creds.Name,
					Phone: users.Phone,
					Opt:   users.Opt,
					Token: CreateToken(w, r, phone),
				}
				InsertUserDB(user)
			}
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("added"))
			return
		} else {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("invalid credinations"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("the user ins't found"))
		// must save it in authoData then check if otp then create token
		opt := CreateOPT()
		auth := db.AuthoData{
			Phone: phone,
			Opt:   opt,
		}
		InsertAutoData(auth)

	}

}
