package functions

import (
	"crypto/rand"
	"demo/db"
	"io"
)

func GetUserDB(phone string) (db.User, bool) {
	var users db.User
	dbb := db.DBConn()
	defer dbb.Close()
	err := dbb.QueryRow("SELECT id, name, phone, opt, token FROM login.user WHERE phone = ?", phone).Scan(&users.Id, &users.Name, &users.Phone, &users.Opt, &users.Token)
	if err != nil {
		return users, false
	}
	return users, true
}

func InsertUserDB(user db.User)bool {
	db := db.DBConn()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO login.user VALUES(?,?,?,?,?)")
	insert.Exec(user.Id, user.Name, user.Phone, user.Opt, user.Token)
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
