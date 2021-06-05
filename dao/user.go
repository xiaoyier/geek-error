package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Passwd      string `json:"passwd"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

const (
	_queryUser  = "SELECT * FROM user WHERE user_name = ?"
	_insertUser = "Insert INTO user (user_name, passwd, email, phone_number) VALUES (?, ?, ?, ?)"
)

func (u *User) Insert() error {

	_, err := db.Exec(_insertUser, u.UserName, u.Passwd, u.Email, u.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "dao.user.Insert: insert user error")
	}
	return nil
}

func QueryUser(userName string) (*User, error) {

	user := new(User)
	err := db.QueryRow(_queryUser, userName).Scan(&user)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "dao.user.QueryUser: query user error")
	}
	return user, nil
}
