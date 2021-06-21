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
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
		return
	}

	phone := creds.Phone
	users, found := GetUserAutho(phone)
	SetPhoneRedis(phone)
	SetOtpRedis(users.Otp)
	fmt.Println(GetOtpRedis(),GetPhoneRedis())
	if found {	
		if creds.Otp == users.Otp {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("accepted"))
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w)
			return
		} else {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("invalid credinations"))
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("the user ins't found"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		otp := CreateOTP()
		auth := db.AuthoData{
			Phone: phone,
			Otp:   otp,
		}
		InsertAutoData(auth)
		return
	}

}
