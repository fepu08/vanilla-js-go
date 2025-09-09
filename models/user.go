package models

type User struct {
	ID             int     `json:"id"`
	Name           int     `json:"name"`
	Email          int     `json:"email"`
	PasswordHashed []byte  `json:"password_hashed"`
	Favorites      []Movie `json:"favorites"`
	Watchlist      []Movie `json:"watch_list"`
}
