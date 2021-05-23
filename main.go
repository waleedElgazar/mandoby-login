package main

import (
	"demo/functions"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("")
	http.HandleFunc("/login", functions.Signin)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

/*
func main() {
	phone:="010100"
	var opt, token,name string
	var id int64

	user,found:=functions.GetUserDB(phone)
	if found {
		opt=user.Opt
		token=user.Token
		name=user.Name
		id=user.Id
		fmt.Println(id,name,phone,opt,token)
	}else{
		user :=db.User{
			Id: 5,
			Name: "waleed",
			Phone: phone,
			Opt: "456123",
			Token: "5456546",
		}
		functions.InsertUserDB(user)
		fmt.Println("created succ")
	}
}
*/
