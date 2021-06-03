package functions

import (
	"crypto/rand"
	"demo/db"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetUserDB(phone string) (db.User, bool) {
	var users db.User
	dbb := db.DBConn()
	defer dbb.Close()
	err := dbb.QueryRow("SELECT name, phone, opt, token FROM login.user WHERE phone = ?", phone).Scan(&users.Name, &users.Phone, &users.Opt, &users.Token)
	if err != nil {
		return users, false
	}
	return users, true
}

func GetUserAutho(phone string) (db.AuthoData, bool) {
	var users db.AuthoData
	dbb := db.DBConn()
	defer dbb.Close()
	err := dbb.QueryRow("SELECT phone, opt FROM login.AuthoData WHERE phone = ?", phone).Scan(&users.Phone, &users.Opt)
	if err != nil {
		return users, false
	}
	return users, true
}

func InsertUserDB(user db.User) bool {
	db := db.DBConn()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO login.user VALUES(?,?,?,?)")
	insert.Exec(user.Name, user.Phone, user.Opt, user.Token)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func InsertAutoData(auth db.AuthoData) bool {
	db := db.DBConn()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO login.AuthoData VALUES(?,?)")
	insert.Exec(auth.Phone, auth.Opt)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func CreateOPT() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
func CreateToken(w http.ResponseWriter, r *http.Request, phone string) string {
	//var creds db.User
	var jwtKey = []byte("c")
	//var user db.User
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := db.Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return tokenString

}
