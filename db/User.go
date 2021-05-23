package db

type User struct {
	Name  string `json:"name"`
	Id    int64 `json:"id"`
	Opt   string `json:"opt"`
	Token string `json:"token"`
	Phone string `json:"phone"`
}


