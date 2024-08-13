package models

import (
	"example.com/events-go/db"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO users (email, password) 
    VALUES (?, ?)
  `)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.Id = id

	return err
}
