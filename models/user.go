package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

//Accout ...
type Accout struct {
	Username string `json:"username"`
	Password string `json:"passphrase"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Role     int    `json:"role"`
}

// NewAcount ...
func NewAcount(acc Accout) *Accout {
	var user Accout
	user.Username = acc.Username
	user.Password = acc.Password
	user.Fullname = acc.Fullname
	user.Email = acc.Email
	user.Role = acc.Role
	return &user
}

// Add ...
func (a *Accout) Add(db *sqlx.DB) (int, error) {
	var id int
	err := db.QueryRow(`insert into users(username, passphrase, fullname, email, role) values ($1, crypt($2, gen_salt('bf', 8)), $3, $4, $5 ) RETURNING id`, a.Username, a.Password, a.Fullname, a.Email, a.Role).Scan(&id)

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code.Name() == "unique_violation" {
				return -1, err
			}
		}
		return 0, err
	}
	return id, nil
}
